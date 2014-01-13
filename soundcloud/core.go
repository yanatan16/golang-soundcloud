// Package soundcloud provides a minimialist soundcloud API wrapper.
package soundcloud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	baseUrl = "https://api.soundcloud.com"
)

type Api struct {
	ClientId    string
	AccessToken string
}

// Create an API with either a ClientId OR an accessToken. Only one is required. Access tokens are preferred because they keep rate limiting down.
func New(clientId string, accessToken string) *Api {
	if clientId == "" && accessToken == "" {
		panic("ClientId or AccessToken must be given to create an Api")
	}

	return &Api{
		ClientId:    clientId,
		AccessToken: accessToken,
	}
}

func (api *Api) Authenticated() bool {
	return api.AccessToken != ""
}

// -- Implementation of request --

func buildGetRequest(urlStr string, params url.Values) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// If we are getting, then we can't merge query params
	if params != nil {
		if u.RawQuery != "" {
			return nil, fmt.Errorf("Cannot merge query params in urlStr and params")
		}
		u.RawQuery = params.Encode()
	}

	u.Path += ".json"
	return http.NewRequest("GET", u.String(), nil)
}

func (api *Api) extendParams(p url.Values, keyvalues ...string) url.Values {
	if p == nil {
		p = url.Values{}
	}
	for i, value := range keyvalues {
		if i % 2 == 0 { continue; }
		p.Set(keyvalues[i-1], value)
	}
	if api.AccessToken != "" {
		p.Set("access_token", api.AccessToken)
	} else {
		p.Set("client_id", api.ClientId)
	}
	return p
}

func (api *Api) get(path string, params url.Values, r interface{}) error {
	params = api.extendParams(params)
	req, err := buildGetRequest(urlify(path), params)
	if err != nil {
		return err
	}
	return api.do(req, r)
}

func (api *Api) do(req *http.Request, r interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return newApiError(req, resp)
	}

	if err := decodeResponse(resp.Body, r); err != nil {
		return newApiErrorStack(err, req, resp)
	}
	return nil
}

func decodeResponse(body io.Reader, to interface{}) error {
	err := json.NewDecoder(body).Decode(to)

	if err != nil {
		return fmt.Errorf("soundcloud: error decoding body: %s", err.Error())
	}
	return nil
}

func urlify(path string) string {
	return baseUrl + path
}

func ensureParams(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}
