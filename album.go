package spotify

import (
	"net/url"
	"strconv"
)

// Album represents an AlbumObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-albumobject
type Album struct {
	Meta
	AlbumType            string    `json:"album_type"`
	AvailableMarkets     []string  `json:"available_markets"`
	Images               []Image   `json:"images"`
	Label                string    `json:"label"`
	Popularity           int       `json:"popularity"`
	ReleaseDate          string    `json:"release_date"`
	ReleaseDatePrecision string    `json:"release_date_precision"`
	TotalTracks          int       `json:"total_tracks"`
	Tracks               TrackPage `json:"tracks"`
	Name                 string    `json:"name"`
}

// SearchAlbum returns a AlbumPaging object that includes a number of albums as per the limit set
func (a *API) SearchAlbum(q, searchType string, limit int) (*PagingAlbums, error) {
	v := url.Values{}
	v.Add("q", q)
	v.Add("type", searchType)
	v.Add("limit", strconv.Itoa(limit))

	pagingAlbums := new(PagingAlbums)
	err := a.get("v1", "/search", v, pagingAlbums)

	return pagingAlbums, err
}

func (a *API) GetAlbum(name string) (*Album, error) {

	pagingAlbums, err := a.SearchAlbum(name, "album", 1)
	if err != nil {
		return nil, err
	}

	album := pagingAlbums.Albums.Items

	return album[0], nil
}
