package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"
)

type Playback struct {
	IsPlaying    bool   `json:"is_playing"`
	Item         Item   `json:"item"`
	ProgressMs   int    `json:"progress_ms"`
	RepeatState  string `json:"repeat_state"`
	ShuffleState bool   `json:"shuffle_state"`
}

type Item struct {
	Artists    []Artist `json:"artists"`
	DurationMs int      `json:"duration_ms"`
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Show       Show     `json:"show"`
	Type       string   `json:"type"`
}

type Artist struct {
	Name string `json:"name"`
}

type Show struct {
	Name string `json:"name"`
}

// GetPlayback gets information about the user's current playback state, including track or episode, progress, and active device.
func (a *API) GetPlayback() (*Playback, error) {
	v := url.Values{}
	v.Add("additional_types", "episode")

	playback := new(Playback)
	err := a.get("/me/player?"+v.Encode(), playback)
	if err == io.EOF {
		return nil, errors.New("Player command failed: No active device found")
	}

	return playback, err
}

// Pause pauses playback on the user's account.
func (a *API) Pause() error {
	return a.put("/me/player/pause", nil)
}

// Play starts a new context or resume current playback on the user's active device.
func (a *API) Play(uris ...string) error {
	if len(uris) == 0 {
		return a.put("/me/player/play", nil)
	}

	type Body struct {
		URIs []string `json:"uris"`
	}

	body := new(Body)
	body.URIs = uris

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return a.put("/me/player/play", bytes.NewReader(data))
}

// Queue adds an item to the end of the user's current playback queue.
func (a *API) Queue(uri string) error {
	v := url.Values{}
	v.Add("uri", uri)

	return a.post("/me/player/queue?"+v.Encode(), nil)
}

// Repeat sets the repeat mode for the user's playback. Options are repeat-track, repeat-context, and off.
func (a *API) Repeat(state string) error {
	v := url.Values{}
	v.Add("state", state)

	return a.put("/me/player/repeat?"+v.Encode(), nil)
}

// Shuffle toggles shuffle on or off for user's playback.
func (a *API) Shuffle(state bool) error {
	v := url.Values{}
	v.Add("state", strconv.FormatBool(state))

	return a.put("/me/player/shuffle?"+v.Encode(), nil)
}

// SkipToPreviousTrack skips to the previous track in the user's queue.
func (a *API) SkipToPreviousTrack() error {
	return a.post("/me/player/previous", nil)
}

// SkipToNextTrack skips to the next track in the user's queue.
func (a *API) SkipToNextTrack() error {
	return a.post("/me/player/next", nil)
}
