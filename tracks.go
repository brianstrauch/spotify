package spotify

import (
	"strconv"
	"time"
)

type Tracks struct {
	ItemsMeta
	Items []Track `json:"items"`
}

type Track struct {
	AddedAt time.Time `json:"added_at"`
	AddedBy Meta      `json:"added_by"`
	IsLocal bool      `json:"is_local"`
	Track   TrackInfo `json:"track"`
	URI     string    `json:"uri"`
}

type TrackInfo struct {
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
