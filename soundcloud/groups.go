package soundcloud

import (
	"net/url"
)

type GroupApi struct {
	groupEndpoint
}

func (api *Api) Groups(params url.Values) ([]Group, error) {
	ret := make([]Group, 0)
	err := api.get("/groups", params, &ret)
	return ret, err
}

func (api *Api) Group(id uint64) *GroupApi {
	return &GroupApi{*api.newGroupEndpoint("groups", id)}
}

func (g *GroupApi) Moderators(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := g.api.get(g.base + "/moderators", params, &ret)
	return ret, err
}

func (g *GroupApi) Members(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := g.api.get(g.base + "/members", params, &ret)
	return ret, err
}

func (g *GroupApi) Contributors(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := g.api.get(g.base + "/contributors", params, &ret)
	return ret, err
}

func (g *GroupApi) Users(params url.Values) ([]User, error) {
	ret := make([]User, 0)
	err := g.api.get(g.base + "/users", params, &ret)
	return ret, err
}

func (g *GroupApi) Tracks(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := g.api.get(g.base + "/tracks", params, &ret)
	return ret, err
}

func (g *GroupApi) PendingTracks(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := g.api.get(g.base + "/pending_tracks", params, &ret)
	return ret, err
}

func (g *GroupApi) PendingTrack(id uint64) (*trackEndpoint) {
	return g.api.newTrackEndpoint(g.base, "pending_tracks", id)
}

func (g *GroupApi) Contributions(params url.Values) ([]Track, error) {
	ret := make([]Track, 0)
	err := g.api.get(g.base + "/contributions", params, &ret)
	return ret, err
}

func (g *GroupApi) Contribution(id uint64) (*trackEndpoint) {
	return g.api.newTrackEndpoint(g.base, "contributions", id)
}
