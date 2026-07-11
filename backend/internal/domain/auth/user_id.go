package auth

type UserID string

func NewUserID(id string) (UserID, error) {
	if id == "" {
		return "", ErrUserIDEmpty
	}

	return UserID(id), nil
}
