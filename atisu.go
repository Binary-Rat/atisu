package atisu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"

	"github.com/pkg/errors"
)

const (
	host     = "api.ati.su"
	trueDemo = "demo=true"
	//enpoints
	getCityID      = "/v1.0/dictionaries/locations/parse"
	searchByFilter = "/v1.0/trucks/search/by-filter"
)

var (
	allowedItemsPerPage = []int{10, 20, 30, 40, 50, 100}
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
	if !slices.Contains(allowedItemsPerPage, itemsPerPage) {
		return nil, errors.Errorf("items per page must be one of %v", allowedItemsPerPage)
	}

	params := map[string]string{}
	if c.isDemo {
		params["demo"] = "true"
	}

	body := requestCars{Page: page, ItemsPerPage: itemsPerPage, Filter: filter}
	return c.doHTTP(context.TODO(), http.MethodGet, searchByFilter, params, body)
}

// GetCityID gets the id of a city by it's name or it`s part`.
// It returns a pointer to a Cities struct (map), which contains the id of the city.
func (c *Client) GetCityID(body []string) (*Cities, error) {
	var cities *Cities
	resp, err := c.doHTTP(context.TODO(), http.MethodPost, getCityID, nil, body)
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
	if len(params) > 0 {
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
func (c *Client) doHTTP(ctx context.Context, method string, path string, params map[string]string, body interface{}) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}
	urlWithHost := endpoint(path, params)
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

	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to read request body")
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("API Error: %s", respB)
		return nil, errors.New(err)
	}

	return respB, nil
}
