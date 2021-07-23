package spotify

import (
	secure "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const accountsBaseURL = "https://accounts.spotify.com"

const (
	// ScopeUGCImageUpload gives write access to user-provided images.
	ScopeUGCImageUpload = "ugc-image-upload"
	// ScopeUserReadRecentlyPlayed gives read access to a user's recently played tracks.
	ScopeUserReadRecentlyPlayed = "user-read-recently-played"
	// ScopeUserTopRead gives read access to a user's top artists and tracks.
	ScopeUserTopRead = "user-top-read"
	// ScopeUserReadPlaybackPosition gives read access to a user's playback position in a content.
	ScopeUserReadPlaybackPosition = "user-read-playback-position"
	// ScopeUserReadPlaybackState gives read access to a user's player state.
	ScopeUserReadPlaybackState = "user-read-playback-state"
	// ScopeUserModifyPlaybackState gives write access to a user's playback state.
	ScopeUserModifyPlaybackState = "user-modify-playback-state"
	// ScopeUserReadCurrentlyPlaying gives read access to a user's currently playing content.
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing"
	// ScopePlaylistModifyPublic gives write access to a user's public playlists.
	ScopePlaylistModifyPublic = "playlist-modify-public"
	// ScopePlaylistModifyPrivate gives write access to a user's private playlists.
	ScopePlaylistModifyPrivate = "playlist-modify-private"
	// ScopePlaylistReadPrivate gives read access to a user's private playlists.
	ScopePlaylistReadPrivate = "playlist-read-private"
	// ScopePlaylistReadCollaborative includes collaborative playlists when requesting a user's playlists.
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	// ScopeUserFollowModify gives write/delete access to the list of artists and other users that the user follows.
	ScopeUserFollowModify = "user-follow-modify"
	// ScopeUserFollowRead gives read access to the list of artists and other users that the user follows.
	ScopeUserFollowRead = "user-follow-read"
	// ScopeUserLibraryModify gives write/delete access to a user's "Your Music" library.
	ScopeUserLibraryModify = "user-library-modify"
	// ScopeUserLibraryRead gives read access to a user's library.
	ScopeUserLibraryRead = "user-library-read"
	// ScopeUserReadEmail gives read access to user's email address.
	ScopeUserReadEmail = "user-read-email"
	// ScopeUserReadPrivate gives read access to user's subscription details (type of user account).
	ScopeUserReadPrivate = "user-read-private"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// CreatePKCEVerifierAndChallenge generates a secure, random code verifier and code challenge for PKCE Authorization.
func CreatePKCEVerifierAndChallenge() (string, string, error) {
	verifier, err := generateRandomVerifier()
	if err != nil {
		return "", "", err
	}

	challenge := calculateChallenge(verifier)

	return string(verifier), challenge, nil
}

func generateRandomVerifier() ([]byte, error) {
	seed, err := generateSecureSeed()
	if err != nil {
		return nil, err
	}
	rand.Seed(seed)

	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_.-~"

	verifier := make([]byte, 128)
	for i := 0; i < len(verifier); i++ {
		idx := rand.Intn(len(chars))
		verifier[i] = chars[idx]
	}

	return verifier, nil
}

func generateSecureSeed() (int64, error) {
	buf := make([]byte, 8)
	_, err := secure.Read(buf)
	if err != nil {
		return 0, err
	}

	seed := int64(binary.BigEndian.Uint64(buf))
	return seed, nil
}

func calculateChallenge(verifier []byte) string {
	hash := sha256.Sum256(verifier)
	challenge := base64.URLEncoding.EncodeToString(hash[:])
	return strings.TrimRight(challenge, "=")
}

// GenerateRandomState creates a random hex string used to mitigate cross-site request forgery attacks.
func GenerateRandomState() (string, error) {
	buf := make([]byte, 7)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	state := hex.EncodeToString(buf)
	return state, nil
}

// BuildAuthURI constructs the URI which users will be redirected to, to authorize the app.
func BuildAuthURI(clientID, redirectURI, state string, showDialog bool, scopes ...string) string {
	q := url.Values{}
	q.Add("client_id", clientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", redirectURI)
	q.Add("code_challenge_method", "S256")
	q.Add("state", state)
	q.Add("scope", strings.Join(scopes, " "))
	q.Add("show_dialog", strconv.FormatBool(showDialog))

	return accountsBaseURL + "/authorize?" + q.Encode()
}

// BuildPKCEAuthURI constructs the URI which users will be redirected to, to authorize the app.
func BuildPKCEAuthURI(clientID, redirectURI, challenge, state string, scopes ...string) string {
	q := url.Values{}
	q.Add("client_id", clientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", redirectURI)
	q.Add("code_challenge_method", "S256")
	q.Add("code_challenge", challenge)
	q.Add("state", state)
	q.Add("scope", strings.Join(scopes, " "))

	return accountsBaseURL + "/authorize?" + q.Encode()
}

// RequestToken allows a user to exchange an authorization code for an access token.
func RequestToken(clientID, code, redirectURI, clientSecret string) (*Token, error) {
	query := make(url.Values)
	query.Set("client_id", clientID)
	query.Set("grant_type", "authorization_code")
	query.Set("code", code)
	query.Set("redirect_uri", redirectURI)
	query.Set("client_secret", clientSecret)
	body := strings.NewReader(query.Encode())

	return postToken(body)
}

// RequestPKCEToken allows a user to exchange an authorization code for an access token.
func RequestPKCEToken(clientID, code, redirectURI, verifier string) (*Token, error) {
	query := make(url.Values)
	query.Set("client_id", clientID)
	query.Set("grant_type", "authorization_code")
	query.Set("code", code)
	query.Set("redirect_uri", redirectURI)
	query.Set("code_verifier", verifier)
	body := strings.NewReader(query.Encode())

	return postToken(body)
}

// RefreshToken allows a user to exchange a refresh token for an access token.
func RefreshToken(refreshToken, clientID, clientSecret string) (*Token, error) {
	query := make(url.Values)
	query.Set("grant_type", "refresh_token")
	query.Set("refresh_token", refreshToken)
	query.Set("client_id", clientID)
	query.Set("client_secret", clientSecret)
	body := strings.NewReader(query.Encode())

	return postToken(body)
}

// RefreshPKCEToken allows a user to exchange a refresh token for an access token.
func RefreshPKCEToken(refreshToken, clientID string) (*Token, error) {
	query := make(url.Values)
	query.Set("grant_type", "refresh_token")
	query.Set("refresh_token", refreshToken)
	query.Set("client_id", clientID)
	body := strings.NewReader(query.Encode())

	return postToken(body)
}

func postToken(body io.Reader) (*Token, error) {
	res, err := http.Post(accountsBaseURL+"/api/token", "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	token := new(Token)
	err = json.NewDecoder(res.Body).Decode(token)

	return token, err
}
