package auth

import (
	"context"
)

type AccessTokenRepository interface {
	GetByToken(ctx context.Context, token string) (AccessToken, error)
	Create(ctx context.Context, token AccessToken) (AccessToken, error)
	Revoke(ctx context.Context, token string) error
}

type AccessTokenService struct {
	repo AccessTokenRepository
}

func NewAccessTokenService(repo AccessTokenRepository) *AccessTokenService {
	return &AccessTokenService{
		repo: repo,
	}
}

func (s *AccessTokenService) GetByToken(ctx context.Context, token string) (AccessToken, error) {
	t, err := s.repo.GetByToken(ctx, token)
	if err != nil {
		return NewZeroAccessToken(), err
	}
	return t, nil
}

func (s *AccessTokenService) Generate(ctx context.Context, userID UserID) (AccessToken, error) {
	newToken, err := NewAccessToken(userID)
	if err != nil {
		return NewZeroAccessToken(), err
	}

	return s.repo.Create(ctx, newToken)
}

func (s *AccessTokenService) Revoke(ctx context.Context, token string) error {
	t, err := s.repo.GetByToken(ctx, token)
	if err != nil {
		return err
	}

	if t.IsRevoked() {
		return nil
	}

	return s.repo.Revoke(ctx, token)
}
