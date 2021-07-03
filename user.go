package spotify

// PublicUser represents a PublicUserObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-publicuserobject
type PublicUser struct {
	Meta
	DisplayName string  `json:"display_name"`
	Images      []Image `json:"images"`
}
