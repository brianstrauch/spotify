package spotify

type User struct {
	Country     string `json:"country"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Product     string `json:"product"`
}

// GetUserProfile gets detailed profile information about the current user (including the current user's username).
func (a *API) GetUserProfile() (*User, error) {
	user := new(User)
	err := a.get("v1", "/me", nil, user)

	return user, err
}
