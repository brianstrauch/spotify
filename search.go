package spotify

import (
	"net/url"
	"strconv"
)

// Search gets Spotify Catalog information about albums, artists, playlists, tracks, shows or episodes that match a
// keyword string.
func (a *API) Search(q, searchType string, limit int) (*Paging, error) {
	query := make(url.Values)
	query.Add("q", q)
	query.Add("type", searchType)
	query.Add("limit", strconv.Itoa(limit))

	paging := new(Paging)
	err := a.get("v1", "/search", query, paging)

	return paging, err
}
