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

// Device represents a DeviceObject in the Spotify API
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

type Item struct {
	Track
	Show Show   `json:"show"`
	Type string `json:"type"`
}

// Show represents a ShowObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-showobject
type Show struct {
	Name string `json:"name"`
}

// GetPlayback gets information about the user's current playback state, including track or episode, progress, and active device.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-information-about-the-users-current-playback
func (a *API) GetPlayback() (*Playback, error) {
	v := url.Values{}
	v.Add("additional_types", "episode")

	playback := new(Playback)
	err := a.get("v1", "/me/player", v, playback)
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
func (a *API) Play(deviceID string, uris ...string) error {
	v := make(url.Values)
	if deviceID != "" {
		v.Add("device_id", deviceID)
	}

	if len(uris) == 0 {
		return a.put("v1", "/me/player/play", v, nil)
	}

	body := &struct {
		URIs []string `json:"uris"`
	}{URIs: uris}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return a.put("v1", "/me/player/play", v, bytes.NewReader(data))
}

// Pause pauses playback on the user's account.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-pause-a-users-playback
func (a *API) Pause(deviceID string) error {
	v := make(url.Values)
	if deviceID != "" {
		v.Add("device_id", deviceID)
	}

	return a.put("v1", "/me/player/pause", v, nil)
}

// SkipToPreviousTrack skips to the previous track in the user's queue.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-skip-users-playback-to-previous-track
func (a *API) SkipToPreviousTrack() error {
	return a.post("v1", "/me/player/previous", nil, nil)
}

// SkipToNextTrack skips to the next track in the user's queue.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-skip-users-playback-to-next-track
func (a *API) SkipToNextTrack() error {
	return a.post("v1", "/me/player/next", nil, nil)
}

// Repeat sets the repeat mode for the user's playback. Options are repeat-track, repeat-context, and off.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-set-repeat-mode-on-users-playback
func (a *API) Repeat(state string) error {
	v := url.Values{}
	v.Add("state", state)

	return a.put("v1", "/me/player/repeat", v, nil)
}

// Shuffle toggles shuffle on or off for user's playback.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-toggle-shuffle-for-users-playback
func (a *API) Shuffle(state bool) error {
	v := url.Values{}
	v.Add("state", strconv.FormatBool(state))

	return a.put("v1", "/me/player/shuffle", v, nil)
}

// Queue adds an item to the end of the user's current playback queue.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-add-to-queue
func (a *API) Queue(uri string) error {
	v := url.Values{}
	v.Add("uri", uri)

	return a.post("v1", "/me/player/queue", v, nil)
}
