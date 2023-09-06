package domain

import "context"

type User struct {
	AutoIncr
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, user User) error
}

type AuthUsecase interface {
	Login(ctx context.Context, user User) (string, error)
	Signup(ctx context.Context, user User) error
	Logout(ctx context.Context, user User) error
}
