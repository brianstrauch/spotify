package spotify

import "time"

// Album represents an AlbumObject in the Spotify API.
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

// Artist represents an ArtistObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-artistobject
type Artist struct {
	Meta
	Genres     []string `json:"genres"`
	Popularity int      `json:"popularity"`
	Name       string   `json:"name"`
}

// Device represents a DeviceObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-deviceobject
type Device struct {
	ID               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

// Error represents an ErrorObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-errorobject
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ExplicitContentSettings represents a ExplicitContentSettingsObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-explicitcontentsettingsobject
type ExplicitContentSettings struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

// Followers represents a FollowersObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-followersobject
type Followers struct {
	HREF  HREF `json:"href"`
	Total int  `json:"total"`
}

// Image represents an ImageObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-imageobject
type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

// Paging represents a PagingObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-pagingobject
type Paging struct {
	Albums AlbumPage `json:"albums"`
	Tracks TrackPage `json:"tracks"`
}

// Playlist represents a PlaylistObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-playlistobject
type Playlist struct {
	Meta
	Collaborative bool              `json:"collaborative"`
	Description   string            `json:"description"`
	Images        []Image           `json:"images"`
	Name          string            `json:"name"`
	Owner         PublicUser        `json:"owner"`
	Public        bool              `json:"public"`
	SnapshotID    string            `json:"snapshot_id"`
	Tracks        PlaylistTrackPage `json:"tracks"`
}

// PlaylistTrack represents a PlaylistTrackObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-playlisttrackobject
type PlaylistTrack struct {
	AddedAt time.Time `json:"added_at"`
	AddedBy Meta      `json:"added_by"`
	IsLocal bool      `json:"is_local"`
	Track   Track     `json:"track"`
	URI     string    `json:"uri"`
}

// PrivateUser represents a PrivateUserObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-privateuserobject
type PrivateUser struct {
	Meta
	Country         string                  `json:"country"`
	DisplayName     string                  `json:"display_name"`
	Email           string                  `json:"email"`
	ExplicitContent ExplicitContentSettings `json:"explicit_content"`
	Followers       Followers               `json:"followers"`
	Images          []*Image                `json:"images"`
	Product         string                  `json:"product"`
}

// PublicUser represents a PublicUserObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-publicuserobject
type PublicUser struct {
	Meta
	DisplayName string  `json:"display_name"`
	Images      []Image `json:"images"`
}

// Show represents a ShowObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-showobject
type Show struct {
	Name string `json:"name"`
}

// Track represents a TrackObject in the Spotify API.
// https://developer.spotify.com/documentation/web-api/reference/#object-trackobject
type Track struct {
	Meta
	Album            Album             `json:"albumomitempty"`
	Artists          []Artist          `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	Duration         *Duration         `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIDs      map[string]string `json:"external_ids"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewURL       string            `json:"preview_url"`
}
