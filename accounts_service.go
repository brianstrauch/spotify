package spotify

import (
	secure "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const (
	AccountsBaseURL = "https://accounts.spotify.com"
	ClientID        = "81dddfee3e8d47d89b7902ba616f3357"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func StartProof() (string, string, error) {
	verifier, err := generateRandomVerifier()
	if err != nil {
		return "", "", err
	}

	hash := sha256.Sum256(verifier)
	challenge := base64.URLEncoding.EncodeToString(hash[:])
	challenge = strings.TrimRight(challenge, "=")

	return string(verifier), challenge, nil
}

func generateRandomVerifier() ([]byte, error) {
	seed, err := generateSecureSeed()
	if err != nil {
		return nil, err
	}
	rand.Seed(seed)

	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_.-~"

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

func BuildAuthURI(redirectURI, challenge, state string, scope string) string {
	q := url.Values{}
	q.Add("client_id", ClientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", redirectURI)
	q.Add("code_challenge_method", "S256")
	q.Add("code_challenge", challenge)
	q.Add("state", state)
	q.Add("scope", scope)

	return AccountsBaseURL + "/authorize?" + q.Encode()
}

func RequestToken(code, redirectURI, verifier string) (*Token, error) {
	v := url.Values{}
	v.Set("client_id", ClientID)
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("redirect_uri", redirectURI)
	v.Set("code_verifier", verifier)
	body := strings.NewReader(v.Encode())

	return postToken(body)
}

func RefreshToken(refreshToken string) (*Token, error) {
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("refresh_token", refreshToken)
	v.Set("client_id", ClientID)
	body := strings.NewReader(v.Encode())

	return postToken(body)
}

func postToken(body io.Reader) (*Token, error) {
	res, err := http.Post(AccountsBaseURL+"/api/token", "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// TODO: Handle errors

	token := new(Token)
	err = json.NewDecoder(res.Body).Decode(token)

	return token, err
}