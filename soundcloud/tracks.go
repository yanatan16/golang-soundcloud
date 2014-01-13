package soundcloud

import (
	"net/url"
)

type TrackApi struct {
	trackEndpoint
}

func (api *Api) Tracks(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := api.get("/tracks", params, &ret)
	return ret, err
}

func (api *Api) Track(id uint64) *TrackApi {
	return &TrackApi{*api.newTrackEndpoint("tracks", id)}
}

func (t *TrackApi) Comments(params url.Values) ([]Comment, error) {
	ret := make([]Comment, 0)
	err := t.api.get(t.base + "/comments", params, &ret)
	return ret, err
}

func (t *TrackApi) Comment(id uint64) (*commentEndpoint) {
	return t.api.newCommentEndpoint(t.base, id)
}