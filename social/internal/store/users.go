package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password, created_at) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	row := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt)

	if err := row.Scan(&user.ID, &user.CreatedAt); err != nil {
		return err
	}
	return nil
}
