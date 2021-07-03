package spotify

import (
	"strconv"
	"time"
)

// Track represents a TrackObject in the Spotify API
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

// Duration is a wrapper struct for time.Duration to unmarshal durations in milliseconds
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
