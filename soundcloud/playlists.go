package soundcloud

import (
	"net/url"
)

type PlaylistApi struct {
	playlistEndpoint
}

func (api *Api) Playlists(params url.Values) ([]Playlist, error) {
	ret := make([]Playlist, 0)
	err := api.get("/playlists", params, &ret)
	return ret, err
}

func (api *Api) Playlist(id uint64) *PlaylistApi {
	return &PlaylistApi{*api.newPlaylistEndpoint("playlists", id)}
}

// No idea how these endpoints works
// func (t *PlaylistApi) SharedToUsers() (*usersEndpoint) {
// func (t *PlaylistApi) SharedToEmails() (*emailsEndpoint) {
// func (t *PlaylistApi) SecretToken() (*tokenEndpoint)

