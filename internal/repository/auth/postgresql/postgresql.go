package auth

import (
	"context"
	"database/sql"
	"main/internal/model"
	def "main/internal/repository"
)

var _ def.AuthRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func New(storagePath string) *repository {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		panic(err.Error())
	}

	return &repository{db: db}
}

func (r *repository) Login(ctx context.Context, email, password string) (string, error) {
	return "", nil
}

func (r *repository) Register(ctx context.Context, userUUID string, info *model.UserInfo) error {
	return nil
}
