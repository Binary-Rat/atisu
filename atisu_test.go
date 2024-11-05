package atisu

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.ati.su",
		Path:   "/v1.0/trucks/search/by-filter",
	}
	u.RawQuery = "demo=true"
	fmt.Println(u.String())
}

func Test_endpoint(t *testing.T) {
	path := "someEndpoint"
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "endpontWithoutParams", args: args{params: map[string]string{}}, want: "https://api.ati.su/someEndpoint"},
		{name: "endpointWithParams", args: args{params: map[string]string{"demo": "true"}}, want: "https://api.ati.su/someEndpoint?demo=true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endpoint(path, tt.args.params); got != tt.want {
				t.Errorf("endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
