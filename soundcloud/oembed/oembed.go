// Package oembed provides structure definitions for the [OEmbed](http://oembed.com) open standard
package oembed

// Response (Section 2.3.4)
type Response struct {
  Type string `json:"type,omitempty"`
  Version float `json:"version,omitempty"`
  Title string `json:"title,omitempty"`
  AuthorName string `json:"author_name,omitempty"`
  AuthorUrl string `json:"author_url,omitempty"`
  ProviderName string `json:"provider_name,omitempty"`
  ProviderUrl string `json:"provider_url,omitempty"`
  CacheAge uint64 `json:"cache_age,omitempty"`
  ThumbnailUrl string `json:"thumbnail_url,omitempty"`

  // Despite what the standard says, may clients treat this as an html height/width parameter and use strings here sometimes like "100%"
  ThumbnailWidth interface{} `json:"thumbnail_width,omitempty"`
  ThumbnailHeight interface{} `json:"thumbnail_height,omitempty"`

  Html string `json:"url,omitempty"`
  Width interface{} `json:"width,omitempty"`
  Height interface{} `json:"height,omitempty"`
}
