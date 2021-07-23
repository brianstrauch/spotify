package spotify

// GetUserProfile gets detailed profile information about the current user (including the current user's username).
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-current-users-profile
func (a *API) GetUserProfile() (*PrivateUser, error) {
	user := new(PrivateUser)
	err := a.get("v1", "/me", nil, user)

	return user, err
}
