package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"main/internal/model"
	"main/internal/storage"
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
func (s *Storage) SaveUser(ctx context.Context, user_id string, email string, passHash []byte, nickname string, elo uint, createdAt, lastLogin time.Time) (int64, error) {
	const op = "storage.postgresql.SaveUser"

	stmt, err := s.db.PrepareContext(ctx, "INSERT INTO users(user_id,email, pass_hash, nickname, elo, created_at, last_login) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, email, passHash)
	if err != nil {
		if pgErr, ok := err.(*pgError); ok && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// User returns user by email.
func (s *Storage) User(ctx context.Context, email string) (model.User, error) {
	const op = "storage.postgresql.User"

	stmt, err := s.db.PrepareContext(ctx, "SELECT id, email, pass_hash FROM users WHERE email = $1")
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user model.User
	var (
		userEmail string
		passHash  []byte
		nickname  string
		elo       uint
	)

	err = row.Scan(&email, &passHash, &nickname, &elo)
	if err != nil {
		// ...
	}

	user.Info = model.UserInfo{
		Email:    userEmail,
		PassHash: passHash,
		Nickname: nickname,
		Elo:      elo,
	}

	err = row.Scan(&user.ID, user.Info, &user.CreatedAt, &user.LastLogin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return model.User{}, fmt.Errorf("%s: %w", op, err)
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
