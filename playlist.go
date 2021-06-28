package spotify

import "path"

// Playlist represents the PlaylistObject struct in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-playlistobject
type Playlist struct {
	Meta
	Collaborative bool       `json:"collaborative"`
	Description   string     `json:"description"`
	Followers     ItemsMeta  `json:"followers"`
	Images        []Image    `json:"images"`
	Name          string     `json:"name"`
	Owner         PublicUser `json:"owner"`
	Public        bool       `json:"public"`
	SnapshotID    string     `json:"snapshot_id"`
	Tracks        Tracks     `json:"tracks"`
}

// PlaylistResponse is the response type for the Get a Playlist request
type PlaylistResponse struct {
	ItemsMeta
	Items []Playlist `json:"items"`
}

func (a *API) GetPlaylists() ([]Playlist, error) {
	playlistsResponse := new(PlaylistResponse)
	if err := a.get("/me/playlists", playlistsResponse); err != nil {
		return nil, err
	}
	// TODO: do more with response
	return playlistsResponse.Items, nil
}

func (a *API) GetPlaylist(playlistID string) (*Playlist, error) {
	playlistResponsse := new(Playlist)
      url := fmt.Sprintf("/playlists/%d", playlistID)
	if err := a.get(url, playlist); err != nil {
		return nil, err
	}
	return playlistResponsse, nil
}
