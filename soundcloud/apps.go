package soundcloud

import (
	"net/url"
)

type AppApi struct {
	appEndpoint
}

func (api *Api) Apps(params url.Values) ([]App, error) {
	ret := make([]App, 0)
	err := api.get("/apps", params, &ret)
	return ret, err
}

func (api *Api) App(id uint64) (*AppApi) {
	return &AppApi{*api.newAppEndpoint("apps", id)}
}

func (a *AppApi) Tracks(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := a.api.get(a.base + "/tracks", params, &ret)
	return ret, err
}