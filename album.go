package spotify

import "time"

// Album represents the AlbumObject type in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-albumobject
type Album struct {
	Meta
	AlbumType            AlbumType `json:"album_type"`
	AvailableMarkets     []string  `json:"available_markets"`
	Images               []Image   `json:"images"`
	Label                string    `json:"label"`
	Popularity           int       `json:"popularity"`
	ReleaseDate          time.Time `json:"release_date"`
	ReleaseDatePrecision Precision `json:"release_date_precision"`
	TotalTracks          int       `json:"total_tracks"`
	Tracks               Tracks    `json:"tracks"`
	Name                 string    `json:"name"`
}

type Precision string

const (
	Day   Precision = "day"
	Month Precision = "month"
	Year  Precision = "year"
)

type AlbumType string

const (
	TypeSingle AlbumType = "single"
	TypeAlbum  AlbumType = "album"
)
