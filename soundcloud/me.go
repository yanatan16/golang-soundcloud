package soundcloud

import (
	"net/url"
)

type MeApi struct {
	UserApi
}

func (api *Api) Me() *MeApi {
	return &MeApi{UserApi{*api.newUserEndpoint("me", true)}}
}

func (me *MeApi) Connections(params url.Values) ([]*User, error) {
	ret := make([]*User, 0)
	err := me.api.get(me.base+"/connections", params, &ret, true)
	return ret, err
}

func (me *MeApi) Connection(id uint64) *userEndpoint {
	return me.api.newUserEndpoint(me.base, "connections", id)
}

func (me *MeApi) Activities(params url.Values) (*PaginatedActivities, error) {
	ret := new(PaginatedActivities)
	err := me.api.get(me.base+"/activities", params, ret, true)
	return ret, err
}
