package soundcloud

import (
)

type CommentApi struct {
	commentEndpoint
}

func (api *Api) Comment(id uint64) *CommentApi {
	return &CommentApi{*api.newCommentEndpoint("comments", id)}
}

// No idea how these endpoints works
// func (t *PlaylistApi) SharedToUsers() (*usersEndpoint) {
// func (t *PlaylistApi) SharedToEmails() (*emailsEndpoint) {
// func (t *PlaylistApi) SecretToken() (*tokenEndpoint)

