package spotify

type Album struct {
	Meta
	AlbumType        AlbumType `json:"album_type"`
	AvailableMarkets []string  `json:"available_markets"`
	Images           []Image   `json:"images"`
	Name             string    `json:"name"`
}

type AlbumType string

const (
	TypeSingle AlbumType = "single"
	TypeAlbum  AlbumType = "album"
)
