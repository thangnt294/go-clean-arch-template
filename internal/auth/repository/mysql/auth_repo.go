package mysql

import (
	"context"
	"go-template/internal/domain"

	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	DB *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) domain.AuthRepository {
	return &authRepo{DB: db}
}

// GetByEmail implements domain.AuthRepository.
func (r *authRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.DB.Get(&user, "SELECT * FROM user WHERE email = ? LIMIT 1", email)
	return user, err
}

// Create implements domain.AuthRepository.
func (r *authRepo) Create(ctx context.Context, user domain.User) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}
