package soundcloud

// User Object (http://developers.soundcloud.com/docs/api/reference#token)
type User struct {
	//---- Included in mini-representation
	Id uint64

	*SubUri
	*SubPermalink

	Username  string
	AvatarUrl string `json:"avatar_url"`

	//---- Not included in mini-representation
	Country              string
	FullName             string `json:"full_name"`
	City                 string
	Description          string
	DiscogsName          string `json:"discogs-name"`
	MyspaceName          string `json:"myspace-name"`
	Website              string
	WebsiteTitle         string `json:"website_title"`
	Online               bool
	TrackCount           uint64 `json:"track_count"`
	PlaylistCount        uint64 `json:"playlist_count"`
	FollowersCount       uint64 `json:"followers_count"`
	FollowingsCount      uint64 `json:"followings_count"`
	PublicFavoritesCount uint64 `json:"public_favorites_count"`
	AvatarData           []byte `json:"avatar_data"`
}

// Sound Track (http://developers.soundcloud.com/docs/api/reference#tracks)
type Track struct {
	Id uint64

	*SubUri
	*SubLabel
	*SubUser
	*SubPermalink

	CreatedAt     string `json:"created_at"`
	Title         string
	Sharing       string
	EmbeddableBy  string `json:"embeddable_by"`
	PurchaseUrl   string `json:"purchase_url"`
	ArtworkUrl    string `json:"artwork_url"`
	Description   string
	Duration      uint64
	Genre         string
	SharedToCount uint64 `json:"shared_to_count"`
	TagList       string `json:"tag_list"`
	// Release             uint64 // Release is sometimes "" and sometimes null and sometimes a number?
	ReleaseDay          uint `json:"release_day"`
	ReleaseMonth        uint `json:"release_month"`
	ReleaseYear         uint `json:"release_year"`
	Streamable          bool
	Downloadable        bool
	State               string
	License             string
	TrackType           string `json:"track_type"`
	WaveformUrl         string `json:"waveform_url"`
	DownloadUrl         string `json:"download_url"`
	StreamUrl           string `json:"stream_url"`
	VideoUrl            string `json:"video_url"`
	Bpm                 float64
	Commentable         bool
	ISRC                string `json:"isrc"`
	KeySignature        string `json:"key_signature"`
	CommentCount        uint64 `json:"comment_count"`
	DownloadCount       uint64 `json:"download_count"`
	PlaybackCount       uint64 `json:"playback_count"`
	FavoritingsCount    uint64 `json:"favoritings_count"`
	OriginalFormat      string `json:"original_format"`
	OriginalContentSize uint64 `json:"original_content_size"`
	CreatedWith         *App
	AssetData           []byte `json:"asset_data"`
	ArtworkData         []byte `json:"artwork_data"`
	UserFavorite        bool   `json:"user_favorite"`
}

// A Soundcloud Set (http://developers.soundcloud.com/docs/api/reference#playlists)
type Playlist struct {
	Id uint64

	*SubUri
	*SubLabel
	*SubUser
	*SubPermalink

	CreatedAt     string `json:"created_at"`
	Title         string
	Sharing       string
	EmbeddableBy  string `json:"embeddable_by"`
	PurchaseUrl   string `json:"purchase_url"`
	ArtworkUrl    string `json:"artwork_url"`
	Description   string
	Duration      uint64
	Genre         string
	SharedToCount uint64 `json:"shared_to_count"`
	TagList       string `json:"tag_list"`
	// Release       uint64 //See release above
	ReleaseDay   uint `json:"release_day"`
	ReleaseMonth uint `json:"release_month"`
	ReleaseYear  uint `json:"release_year"`
	Streamable   bool
	Downloadable bool
	EAN          string `json:"ean"`
	PlaylistType string `json:"playlist_type"`
	Tracks       []*Track
}

// Groups of members with tracks (http://developers.soundcloud.com/docs/api/reference#groups)
type Group struct {
	Id uint64

	*SubPermalink

	*SubUri
	CreatedAt        string `json:"created_at"`
	ArtworkUrl       string `json:"artwork_url"`
	Name             string
	Description      string
	ShortDescription string `json:"short_description"`

	// A mini representation of user
	Creator *User
}

// Comment on a track (http://developers.soundcloud.com/docs/api/reference#comments)
type Comment struct {
	Id uint64

	*SubUri
	*SubUser

	CreatedAt string `json:"created_at"`
	Body      string
	// Time in milliseconds
	Timestamp uint64
	TrackId   uint64
}

// A connection to an external service such as twitter, tumblr, facebook, etc (http://developers.soundcloud.com/docs/api/reference#connections)
type Connection struct {
	Id uint64

	*SubUri

	CreatedAt    string `json:"created_at"`
	DisplayName  string `json:"display_name"`
	PostFavorite bool   `json:"post_favorite"`
	PostPublish  bool   `json:"post_publish"`
	Service      string
	Type         string
}

// Activity by a user (http://developers.soundcloud.com/docs/api/reference#activities)
type Activity struct {
	Type      string
	CreatedAt string `json:"created_at"`
	Tags      string
	// This is a specific type that should be re-interpreted as a given type (see activities page)
	Origin interface{}
}

type App struct {
	Id   uint64
	Name string

	*SubUri
	*SubPermalink

	ExternalUrl string `json:"external_url"`
	Creator     *User
}

// -- Special Response objects

type PaginatedActivities struct {
	*Pagination
	Collection []Activity
}

// -- subobjects

type SubUri struct {
	Uri string
}

type SubPermalink struct {
	// The name of the permalink. Sometimes not included (on App specifically)
	Permalink string
	// Always included
	PermalinkUrl string `json:"permalink_url"`
}

type SubLabel struct {
	Label     *User
	LabelId   uint64 `json:"label_id"`
	LabelName string `json:"label_name"`
}

type SubUser struct {
	UserId uint64 `json:"user_id"`
	// A minimal representation of user
	User *User
}

type Pagination struct {
	NextHref string `json:"next_href"`
}

type Resolution struct {
	Status   string
	Location string
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string
}
