package spotify

// Artist represents an ArtistObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-artistobject
type Artist struct {
	Meta
	Genres     []string `json:"genres"`
	Popularity int      `json:"popularity"`
	Name       string   `json:"name"`
}
