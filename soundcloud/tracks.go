package soundcloud

import (
	"net/url"
)

type TrackApi struct {
	trackEndpoint
}

func (api *Api) Tracks(params url.Values) ([]*Track, error) {
	ret := make([]*Track, 0)
	err := api.get("/tracks", params, &ret)
	return ret, err
}

func (api *Api) Track(id uint64) *TrackApi {
	return &TrackApi{*api.newTrackEndpoint("tracks", id)}
}

func (t *TrackApi) Comments(params url.Values) ([]*Comment, error) {
	ret := make([]*Comment, 0)
	err := t.api.get(t.base+"/comments", params, &ret)
	return ret, err
}

func (t *TrackApi) Comment(id uint64) *commentEndpoint {
	return t.api.newCommentEndpoint(t.base, "comments", id)
}

func (t *TrackApi) Favorites(params url.Values) ([]*User, error) {
	ret := make([]*User, 0)
	err := t.api.get(t.base+"/favorites", params, &ret)
	return ret, err
}

func (t *TrackApi) Favorite(id uint64) *userEndpoint {
	return t.api.newUserEndpoint(t.base, "favorites", id)
}

// No idea how these endpoints works
// func (t *TrackApi) SharedToUsers() (*usersEndpoint) {
// func (t *TrackApi) SharedToEmails() (*emailsEndpoint) {
// func (t *TrackApi) SecretToken() (*tokenEndpoint)
