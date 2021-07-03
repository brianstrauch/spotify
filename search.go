package spotify

import (
	"net/url"
	"strconv"
)

// Search gets Spotify Catalog information about albums, artists, playlists, tracks, shows or episodes that match a
// keyword string.
func (a *API) Search(q string, limit int) (*Paging, error) {
	v := url.Values{}
	v.Add("q", q)
	v.Add("type", "track")
	v.Add("limit", strconv.Itoa(limit))

	paging := new(Paging)
	err := a.get("v1", "/search", v, paging)

	return paging, err
}
