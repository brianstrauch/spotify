package spotify

import (
	"strconv"
	"time"
)

// Tracks represents a list of tracks
type Tracks struct {
	ItemsMeta
	Items []PlaylistTrack `json:"items"`
}

// Track represents a Track in the Spotify API
// See https://developer.spotify.com/documentation/web-api/reference/#object-playlisttrackobject
type PlaylistTrack struct {
	// AddedAt is when the track was added to the playlist or saved
	AddedAt time.Time `json:"added_at"`
	// AddedBy represents the user that added the track to the playlist, or saved it
	AddedBy Meta `json:"added_by"`
	IsLocal bool `json:"is_local"`
	// Track contains details about the track itself
	Track Track `json:"track"`
	// URI is the Spotify URI of the track
	URI string `json:"uri"`
}

// TrackObject represents the TrackObject struct in the API
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

// Duration is a wrapper struct for time.Duration to unmarshal durations in miliseconds
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	ms, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	d.Duration = time.Duration(ms * 1000000)
	return nil
}
