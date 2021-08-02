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
	Track
	Show Show `json:"show"`
}

// GetPlayback gets information about the user's current playback state, including track or episode, progress, and active device.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-information-about-the-users-current-playback
func (a *API) GetPlayback() (*Playback, error) {
	query := make(url.Values)
	query.Add("additional_types", "episode")

	playback := new(Playback)
	err := a.get("v1", "/me/player", query, playback)
	if err == io.EOF {
		return nil, errors.New("Player command failed: No active device found")
	}

	return playback, err
}

// GetDevices gets information about a user's available devices.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-users-available-devices
func (a *API) GetDevices() ([]*Device, error) {
	res := &struct {
		Devices []*Device `json:"devices"`
	}{}

	err := a.get("v1", "/me/player/devices", nil, res)
	return res.Devices, err
}

// Play starts a new context or resume current playback on the user's active device.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-start-a-users-playback
func (a *API) Play(contextURI string, uris ...string) error {
	query := make(url.Values)

	body := &struct {
		ContextURIs string   `json:"context_uri,omitempty"`
		URIs        []string `json:"uris,omitempty"`
	}{
		ContextURIs: contextURI,
		URIs:        uris,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return a.put("v1", "/me/player/play", query, bytes.NewReader(data))
}

// Pause pauses playback on the user's account.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-pause-a-users-playback
func (a *API) Pause() error {
	return a.put("v1", "/me/player/pause", nil, nil)
}

// SkipToPreviousTrack skips to the previous track in the user's queue.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-skip-users-playback-to-previous-track
func (a *API) SkipToPreviousTrack() error {
	return a.post("v1", "/me/player/previous", nil, nil, nil)
}

// SkipToNextTrack skips to the next track in the user's queue.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-skip-users-playback-to-next-track
func (a *API) SkipToNextTrack() error {
	return a.post("v1", "/me/player/next", nil, nil, nil)
}

// Repeat sets the repeat mode for the user's playback. Options are repeat-track, repeat-context, and off.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-set-repeat-mode-on-users-playback
func (a *API) Repeat(state string) error {
	query := make(url.Values)
	query.Add("state", state)

	return a.put("v1", "/me/player/repeat", query, nil)
}

// Shuffle toggles shuffle on or off for user's playback.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-toggle-shuffle-for-users-playback
func (a *API) Shuffle(state bool) error {
	query := make(url.Values)
	query.Add("state", strconv.FormatBool(state))

	return a.put("v1", "/me/player/shuffle", query, nil)
}

// Queue adds an item to the end of the user's current playback queue.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-add-to-queue
func (a *API) Queue(uri string) error {
	query := make(url.Values)
	query.Add("uri", uri)

	return a.post("v1", "/me/player/queue", query, nil, nil)
}
