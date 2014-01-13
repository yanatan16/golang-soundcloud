package soundcloud

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUsers(t *testing.T) {
	ret, err := api.Users(url.Values{"q": []string{"ladygaga"}})
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("query returned no values")
	} else if ret[0].Username != "ladygaga" {
		t.Error("lady gaga wasn't the first response?!")
	}
}

func TestUser(t *testing.T) {
	ret, err := api.User(ladygaga_id).Get(nil)
	if err != nil {
		t.Error(err)
	} else if ret.Username != "ladygaga" {
		t.Error("lady gaga's id changed?")
	} else if ret.Permalink != "ladygaga" || ret.PermalinkUrl != "http://soundcloud.com/ladygaga" {
		t.Error("permalink error", ret)
	} else if ret.Uri != fmt.Sprintf("https://api.soundcloud.com/users/%d", ladygaga_id) {
		t.Error("uri", ret)
	} else if ret.Country != "United States" || ret.FullName != "Lady Gaga" || ret.City != "New York City" ||
		ret.Website != "http://littlemonsters.com" || ret.WebsiteTitle != "LittleMonsters.com" || ret.FollowersCount < 1000 {

		t.Error("Something wrong with ladygaga's profile", ret)
	}
}

func TestUserTracks(t *testing.T) {
	ret, err := api.User(ladygaga_id).Tracks(nil)
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("lady gaga has no tracks?")
	} else {
		for _, tr := range ret {
			if tr.UserId != ladygaga_id || tr.User.Id != ladygaga_id {
				t.Error("lady gaga's track isn't owned by her?")
			} else if tr.CreatedAt == "" {
				t.Error("a track doesn't have a createdAt?")
			}
		}
	}
}

func TestUserTracksCreatedAtFrom(t *testing.T) {
	ret, err := api.User(macklemore_id).Tracks(api.Values("created_at[from]", "2013/03/17 18:37:51"))
	if err != nil {
		t.Error(err)
	} else if ret[len(ret) - 1].CreatedAt != "2013/03/17 18:37:51 +0000" {
		t.Error("Created At didn't work right")
	}
}

func TestUserPlaylists(t *testing.T) {
	ret, err := api.User(ladygaga_id).Playlists(nil)
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("no playlists?")
	} else {
		for _, pl := range ret {
			if pl.UserId != ladygaga_id || pl.User.Id != ladygaga_id {
				t.Error("playlist has wrong user")
			}
		}
	}
}

func TestUserFollowings(t *testing.T) {
	_, err := api.User(joneisen_id).Followings(nil)
	if err != nil {
		t.Error(err)
		// } else if len(ret) != 1 {
		// 	t.Error("should be following only one user", ret)
		// } else if ret[0].Id != macklemore_id {
		// 	t.Error("Should be following macklemore")
	}
}

func TestUserFollowers(t *testing.T) {
	ret, err := api.User(ladygaga_id).Followers(nil)
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("should be followed by millions!")
	}
}

func TestUserComments(t *testing.T) {
	_, err := api.User(ladygaga_id).Comments(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUserFavorites(t *testing.T) {
	_, err := api.User(ladygaga_id).Favorites(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUserGroups(t *testing.T) {
	_, err := api.User(ladygaga_id).Groups(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUserWebProfiles(t *testing.T) {
	_, err := api.User(ladygaga_id).WebProfiles(nil)
	if err != nil {
		t.Error(err)
	}
}
