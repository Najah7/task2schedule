package application

import (
	"context"

	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/Najah7/task2schedule/internal/domain/auth"
	"github.com/Najah7/task2schedule/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type config struct {
}

type logger interface{}

type Store struct {
	Users        *repositories.UserRepository
	AccessTokens *repositories.AccessTokenRepository
}

func NewStore(ctx context.Context) Store {
	cfg, err := pgxpool.ParseConfig(postgresConnectionString())
	if err != nil {
		panic(fmt.Errorf("failed to parse connection string: %w", err))
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create connection pool: %w", err))
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		panic(fmt.Errorf("ping postgres: %w", err))
	}

	userRepo := repositories.NewUserRepository(pool)
	accessTokenRepo := repositories.NewAccessTokenRepository(pool)

	return Store{
		Users:        userRepo,
		AccessTokens: accessTokenRepo,
	}
}

func postgresConnectionString() string {
	postgresURL := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(
			getenv("POSTGRES_USER", "changeme"),
			getenv("POSTGRES_PASSWORD", "changeme"),
		),
		Host: net.JoinHostPort(
			getenv("POSTGRES_HOST", "localhost"),
			getenv("POSTGRES_PORT", "5432"),
		),
		Path: "/" + getenv("POSTGRES_DB", "task2schedule"),
	}

	query := postgresURL.Query()
	query.Set("sslmode", getenv("POSTGRES_SSLMODE", "disable"))
	postgresURL.RawQuery = query.Encode()

	return postgresURL.String()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

type Service struct {
	User        *auth.UserService
	AccessToken *auth.AccessTokenService
}

func NewService(store Store) Service {
	userService := auth.NewUserService(store.Users)
	accessTokenService := auth.NewAccessTokenService(store.AccessTokens)

	return Service{
		User:        userService,
		AccessToken: accessTokenService,
	}
}

type Application struct {
	config  config
	logger  logger
	Store   Store
	Service Service
}

func New() *Application {
	ctx := context.Background()
	store := NewStore(ctx)
	service := NewService(store)

	return &Application{
		config:  config{},
		logger:  nil,
		Store:   store,
		Service: service,
	}
}
