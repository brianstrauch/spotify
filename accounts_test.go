package spotify

import (
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func TestCreatePKCEVerifierAndChallenge(t *testing.T) {
	verifier, challenge, err := CreatePKCEVerifierAndChallenge()
	if err != nil {
		t.Fatal(err)
	}

	if !regexp.MustCompile(`^[[:alnum:]_.\-~]{128}$`).MatchString(verifier) {
		t.Fatal("Verifier string does not match")
	}

	// Hash with SHA-256 (64 chars)
	// Convert to Base64 (44 chars)
	// Remove trailing = (43 chars)
	if !regexp.MustCompile(`^[[:alnum:]\-_]{43}$`).MatchString(challenge) {
		t.Fatal("Challenge string does not match")
	}
}

func TestBuildPKCEAuthURI(t *testing.T) {
	var (
		clientID    = "client"
		redirectURI = "http://localhost:1024"
		challenge   = "challenge"
		state       = "state"
		scope       = "user-modify-playback-state"
	)

	uri := BuildPKCEAuthURI(clientID, redirectURI, challenge, state, scope)

	substrings := []string{
		"client_id=" + clientID,
		"response_type=code",
		"redirect_uri=" + url.QueryEscape(redirectURI),
		"code_challenge_method=S256",
		"code_challenge=" + challenge,
		"state=" + state,
		"scope=" + url.QueryEscape(scope),
	}

	for _, substring := range substrings {
		if !strings.Contains(uri, substring) {
			t.Fatalf("URI %s does not contain substring %s", uri, substring)
		}
	}
}
