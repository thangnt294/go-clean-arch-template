package domain

import "context"

type User struct {
	AutoIncr
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email,required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, user User) error
}

type AuthUsecase interface {
	Login(ctx context.Context, user UserLogin) (string, error)
	Signup(ctx context.Context, user User) error
}
