package spotify

import (
	"net/url"
	"strconv"
)

type Paging struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	Items []Track `json:"items"`
}

type Track struct {
	URI string `json:"uri"`
}

// Get Spotify Catalog information about albums, artists, playlists, tracks, shows or episodes that match a keyword
// string.
func (a *API) Search(q string, limit int) (*Paging, error) {
	v := url.Values{}
	v.Add("q", q)
	v.Add("type", "track")
	v.Add("limit", strconv.Itoa(limit))

	pagingObject := new(Paging)
	err := a.get("/search?"+v.Encode(), pagingObject)

	return pagingObject, err
}
