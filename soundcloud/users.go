package soundcloud

import (
	"net/url"
)

type UserApi struct {
	userEndpoint
}

func (api *Api) Users(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := api.get("/users", params, &ret)
	return ret, err
}

func (api *Api) User(id uint64) *UserApi {
	return &UserApi{*api.newUserEndpoint("users", id)}
}

func (u *UserApi) Get(params url.Values) (*User, error) {
	ret := new(User)
	err := u.api.get(u.base, params, ret)
	return ret, err
}

func (u *UserApi) Tracks(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := u.api.get(u.base+"/tracks", params, &ret)
	return ret, err
}

func (u *UserApi) Playlists(params url.Values) ([]Playlist, error) {
	ret := make([]Playlist, 0)
	err := u.api.get(u.base+"/playlists", params, &ret)
	return ret, err
}

func (u *UserApi) Followings(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := u.api.get(u.base+"/followings", params, &ret)
	return ret, err
}

func (u *UserApi) Following(id uint64) *userEndpoint {
	return u.api.newUserEndpoint(u.base, "followings", id)
}

func (u *UserApi) Followers(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := u.api.get(u.base+"/followers", params, &ret)
	return ret, err
}

func (u *UserApi) Follower(id string) *userEndpoint {
	return u.api.newUserEndpoint(u.base, "followers", id)
}

func (u *UserApi) Comments(params url.Values) ([]Comment, error) {
	ret := make([]Comment, 0)
	err := u.api.get(u.base+"/comments", params, &ret)
	return ret, err
}

func (u *UserApi) Favorites(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := u.api.get(u.base+"/favorites", params, &ret)
	return ret, err
}

func (u *UserApi) Favorite(id string) *trackEndpoint {
	return u.api.newTrackEndpoint(u.base, "favorites", id)
}

func (u *UserApi) Groups(params url.Values) ([]Group, error) {
	ret := make([]Group, 0)
	err := u.api.get(u.base+"/groups", params, &ret)
	return ret, err
}

func (u *UserApi) WebProfiles(params url.Values) ([]Connection, error) {
	ret := make([]Connection, 0)
	err := u.api.get(u.base+"/web-profiles", params, &ret)
	return ret, err
}
