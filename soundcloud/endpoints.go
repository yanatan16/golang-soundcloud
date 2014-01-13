package soundcloud

import (
	"net/url"
	"fmt"
)

type endpoint struct {
	api *Api
	base string
}

func (api *Api) newEndpoint(dirs ...interface{}) endpoint {
	path := ""
	for _, dir := range dirs {
		if sdir, ok := dir.(string); ok {
			path += "/" + sdir
			continue
		}
		if idir, ok := dir.(uint64); ok {
			path += fmt.Sprintf("/%d", idir)
			continue
		}
		panic("Only strings and uint64 accepted")
	}
	return endpoint{api, path}
}

type userEndpoint struct {
	endpoint
}

func (api *Api) newUserEndpoint(dirs ...interface{}) *userEndpoint {
	return &userEndpoint{api.newEndpoint(dirs...)}
}

func (u *userEndpoint) Get(params url.Values) (*User, error) {
	ret := new(User)
	err := u.api.get(u.base, params, ret)
	return ret, err
}

// func (u *userEndpoint) Put(params url.Values) (*User, error) {
//   ret := new(User)
//   err := u.api.put(u.base, params, ret)
//   return ret, err
// }
// func (u *userEndpoint) Delete(params url.Values) (*User, error) {
//   ret := new(User)
//   err := u.api.delete(u.base, params, ret)
//   return ret, err
// }

type trackEndpoint struct {
	endpoint
}

func (api *Api) newTrackEndpoint(dirs ...interface{}) *trackEndpoint {
	return &trackEndpoint{api.newEndpoint(dirs...)}
}

func (t *trackEndpoint) Get(params url.Values) (*Track, error) {
	ret := new(Track)
	err := t.api.get(t.base, params, ret)
	return ret, err
}

type commentEndpoint struct {
	endpoint
}

func (api *Api) newCommentEndpoint(dirs ...interface{}) *commentEndpoint {
	return &commentEndpoint{api.newEndpoint(dirs...)}
}

func (t *commentEndpoint) Get(params url.Values) (*Comment, error) {
	ret := new(Comment)
	err := t.api.get(t.base, params, ret)
	return ret, err
}