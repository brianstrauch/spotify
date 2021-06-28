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
	playlists := new(PlaylistResponse)
	if err := a.get("v1", "/me/playlists", nil, playlists); err != nil {
		return nil, err
	}
	return playlists.Items, nil
}

func (a *API) GetPlaylist(id string) (*Playlist, error) {
	playlist := new(Playlist)
	if err := a.get("v1", path.Join("/playlists", id), nil, playlist); err != nil {
		return nil, err
	}
	return playlist, nil
}
