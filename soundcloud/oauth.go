package soundcloud

import (
	"fmt"
)

func (api *Api) Refresh() error {
	if api.ClientId == "" || api.ClientSecret == "" || api.RefreshToken == "" {
		return fmt.Errorf("ClientId, ClientSecret, and RefreshToken must all be specified")
	}

	values := Values(
		"client_id", api.ClientId,
		"client_secret", api.ClientSecret,
		"refresh_token", api.RefreshToken,
		"grant_type", "refresh_token",
	)

	ret := new(AuthResponse)
	err := api.post("/oauth2/token", values, ret)

	if err == nil {
		api.AccessToken = ret.AccessToken
		api.RefreshToken = ret.RefreshToken
	}

	return err
}
