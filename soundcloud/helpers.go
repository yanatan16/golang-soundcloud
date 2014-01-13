package soundcloud

import (
	"net/url"
	"strings"
)


func Values(keyValues ...string) url.Values {
  v := url.Values{}
  for i := 0; i < len(keyValues)-1; i += 2 {
    v.Set(keyValues[i], keyValues[i+1])
  }
  return v
}
func (api *Api) Values(keyValues ...string) url.Values {
  return Values(keyValues...)
}


type userEndpoint struct {
	api     *Api
	base    string
	getOnly bool
}

func (api *Api) newUserEndpoint(getOnly bool, dirs ...string) *userEndpoint {
	return &userEndpoint{api, "/" + strings.Join(dirs, "/"), getOnly}
}

func (u *userEndpoint) Get(params url.Values) (error, *User) {
	ret := new(User)
	err := u.api.get(u.base, params, ret)
	return err, ret
}

// func (u *userEndpoint) Put(params url.Values) (error, *User) {
//   ret := new(User)
//   err := u.api.put(u.base, params, ret)
//   return err, ret
// }
// func (u *userEndpoint) Delete(params url.Values) (error, *User) {
//   ret := new(User)
//   err := u.api.delete(u.base, params, ret)
//   return err, ret
// }


type trackEndpoint struct {
  api     *Api
  base    string
  getOnly bool
}

func (api *Api) newTrackEndpoint(getOnly bool, dirs ...string) *trackEndpoint {
  return &trackEndpoint{api, "/" + strings.Join(dirs, "/"), getOnly}
}

func (t *trackEndpoint) Get(params url.Values) (error, *Track) {
  ret := new(Track)
  err := t.api.get(t.base, params, ret)
  return err, ret
}