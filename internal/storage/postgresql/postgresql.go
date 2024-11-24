package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"main/internal/model"
	"main/internal/storage"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) (*Storage, error) {
	const op = "storage.postgresql.New"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

// SaveUser saves user to db.
func (s *Storage) SaveUser(ctx context.Context, user_id uuid.UUID, email string, passHash []byte, nickname string, elo uint, createdAt, lastLogin time.Time) (uuid.UUID, error) {
	const op = "storage.postgresql.SaveUser"

	stmt, err := s.db.PrepareContext(ctx, "INSERT INTO users(user_id,email, pass_hash, nickname, elo, created_at, last_login) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING user_id")
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	var id uuid.UUID
	err = stmt.QueryRowContext(ctx, user_id, email, passHash, nickname, elo, createdAt, lastLogin).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgError); ok && pgErr.Code == "23505" {
			return uuid.Nil, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// User returns user by email.
func (s *Storage) User(ctx context.Context, email string) (model.User, error) {
	const op = "storage.postgresql.User"

	stmt, err := s.db.PrepareContext(ctx, "SELECT user_id, email, pass_hash, nickname, elo, created_at, last_login FROM users WHERE email = $1")
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	var user model.User
	var (
		userEmail string
		passHash  []byte
		nickname  string
		elo       uint
	)

	err = stmt.QueryRowContext(ctx, email).Scan(&user.ID, &userEmail, &passHash, &nickname, &elo, &user.CreatedAt, &user.LastLogin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user.Info = model.UserInfo{
		Email:    userEmail,
		PassHash: passHash,
		Nickname: nickname,
		Elo:      elo,
	}

	return user, nil
}

type pgError struct {
	Code    string
	Message string
}

func (e *pgError) Error() string {
	return e.Message
}
