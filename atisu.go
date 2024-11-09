package atisu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	host = "api.ati.su"
	//enpoints
	searchByFilter = "/v1.0/trucks/search/by-filter"
	//errors
	xErrorHeader = "X-Error"
)

//go:generate mockgen -destination=./mocks/mock_atisu.go -package=mocks github.com/Binary-Rat/atisu HTTPClient
type Client struct {
	isDemo bool
	client HTTPClient
	token  string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient returns a new Client instance.
// token should be a string, provided by ati.su.
// isDemo is a boolean flag, indicating whether to use the api in demo mode or not.
func NewClient(token string, isDemo bool) (*Client, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	return &Client{
		isDemo: isDemo,
		token:  token,
		client: &http.Client{},
	}, nil
}

func (c *Client) GetCarsWithFilter(page int, itemsPerPage int, filter Filter) ([]byte, error) {
	params := map[string]string{}
	if c.isDemo {
		params["demo"] = "true"
	}
	enpoint := endpoint(searchByFilter, params)
	body := requestCars{Page: page, ItemsPerPage: itemsPerPage, Filter: filter}
	return c.doHTTP(context.TODO(), http.MethodGet, enpoint, body)
}

// Get city id from api ati.su, can`t be cashed in your service to increase performance`
func (c *Client) GetCityID(body []string) ([]City, error) {
	var cities []City
	resp, err := c.doHTTP(context.TODO(), http.MethodGet, getCityID, body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get city id")
	}
	json.Unmarshal(resp, &cities)

	return cities, nil
}

// Note that Filter recive the id of the city not the name
// That why you need to use GetCityID before search by filter

func endpoint(path string, params map[string]string) string {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
	if params == nil || len(params) > 0 {
		q := u.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	return u.String()
}

// The function returns the response body and an error if something went wrong.
// all api responses`s errors will be returned as error
func (c *Client) doHTTP(ctx context.Context, method string, endpoint string, body interface{}) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}
	urlWithHost, err := url.JoinPath(host, endpoint)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to join path")
	}
	req, err := http.NewRequestWithContext(ctx, method, urlWithHost, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create new request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to send http request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("API Error: %s", resp.Header.Get(xErrorHeader))
		return nil, errors.New(err)
	}

	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to read request body")
	}

	return respB, nil
}
