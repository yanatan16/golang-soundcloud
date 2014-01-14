package soundcloud

import (
	"fmt"
	"net/url"
)

func (api *Api) Refresh() error {
	if api.ClientId == "" || api.ClientSecret == "" || api.RefreshToken == "" {
		return fmt.Errorf("ClientId, ClientSecret, and RefreshToken must all be specified")
	}

	values := url.Values{
		"client_id":     []string{api.ClientId},
		"client_secret": []string{api.ClientSecret},
		"refresh_token": []string{api.RefreshToken},
		"grant_type":    []string{"refresh_token"},
	}

	ret := new(AuthResponse)
	err := api.post("/oauth2/token", values, ret)

	if err == nil {
		api.AccessToken = ret.AccessToken
		api.RefreshToken = ret.RefreshToken
	}

	return err
}
