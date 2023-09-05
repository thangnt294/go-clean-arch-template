package usecase

import (
	"context"
	"errors"
	"go-template/config"
	"go-template/internal/auth/util"
	"go-template/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUsecase(authRepo domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{authRepo}
}

func (u *authUsecase) Login(ctx context.Context, user domain.User) (string, error) {
	existingUser, err := u.authRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if !util.CompareHashPassword(existingUser.Password, user.Password) {
		return "", errors.New("Invalid password")
	}

	// TODO: add expire time
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"username": existingUser.Name,
			"email":    existingUser.Email,
			"role":     existingUser.Role,
		})

	return token.SignedString(config.LoadConfig().JWTKey)
}

func (u *authUsecase) Signup(ctx context.Context, user domain.User) error {
	// TODO: implement
	return nil
}

func (u *authUsecase) Logout(ctx context.Context, user domain.User) error {
	// TODO: implement
	return nil
}
