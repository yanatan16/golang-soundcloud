package soundcloud

import (
  "net/url"
  "fmt"
)

type TrackApi struct {
	api  *Api
	base string
}

func (api *Api) Tracks(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := api.get("/tracks", params, &ret)
	return ret, err
}

func (api *Api) Track(id uint64) *TrackApi {
	return &TrackApi{api, fmt.Sprintf("/tracks/%d", id)}
}

func (u *TrackApi) Get(params url.Values) (*Track, error) {
	ret := new(Track)
	err := u.api.get(u.base, params, ret)
	return ret, err
}