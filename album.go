package spotify

type Album struct {
	Meta
	AlbumType        AlbumType `json:"album_type"`
	AvailableMarkets []string  `json:"available_markets"`
	Images           []Image
	Name             string
}

type AlbumType string

const (
	Single_T AlbumType = "single"
	Album_T  AlbumType = "album"
)
