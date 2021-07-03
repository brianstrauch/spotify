package spotify

// Image represents an ImageObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-imageobject
type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
