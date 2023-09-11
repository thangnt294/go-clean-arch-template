package domain

import "context"

type User struct {
	AutoIncr
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, user User) error
}

type AuthUsecase interface {
	Login(ctx context.Context, user User) (string, error)
	Signup(ctx context.Context, user User) error
}
