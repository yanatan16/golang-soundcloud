package soundcloud

import (
	"net/url"
	"testing"
)

func TestTracks(t *testing.T) {
	ret, err := api.Tracks(url.Values{"q": []string{"welcome to night vale"}})
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("query returned no values")
	} else if ret[0].User.Username != "planetoffinks" {
		t.Error("planetoffinks wasn't the creator first response?!")
	}
}

func TestTrackGet(t *testing.T) {
	ret, err := api.Track(thriftshop_id).Get(nil)
	if err != nil {
		t.Error(err)
	} else if ret.Id != thriftshop_id {
		t.Error("id didn't come back as requested")
	} else if ret.Permalink != "macklemore-x-ryan-lewis-thrift" {
		t.Error("thrift shop's permalink changed?")
	} else if ret.Title != "Macklemore X Ryan Lewis - Thrift Shop feat. Wanz" {
		t.Error("thrift shop title changed?")
	} else if ret.UserId != macklemore_id || ret.UserId != ret.User.Id {
		t.Error("user object is messed up", ret.UserId, ret.User)
	} else if ret.User.Permalink != "macklemore" || ret.User.Username != "Macklemore & Ryan Lewis" {
		t.Error("user object for macklmore changed?")
	} else if ret.PlaybackCount < 1000 || ret.FavoritingsCount < 1000 || ret.CommentCount < 1000 {
		t.Error("Some counts are wrong?", ret)
	}
}
