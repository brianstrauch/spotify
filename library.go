package spotify

import (
	"net/url"
	"strings"
)

// SaveTracks saves one or more tracks to the current user's 'Your Music' library.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-save-tracks-user
func (a *API) SaveTracks(ids ...string) error {
	query := make(url.Values)
	query.Add("ids", strings.Join(ids, ","))

	return a.put("v1", "/me/tracks", query, nil)
}

// RemoveSavedTracks removes one or more tracks from the current user's 'Your Music' library.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-remove-tracks-user
func (a *API) RemoveSavedTracks(ids ...string) error {
	query := make(url.Values)
	query.Add("ids", strings.Join(ids, ","))

	return a.delete("v1", "/me/tracks", query)
}
