package usecase

import (
	"context"
	"errors"
	"time"

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

func (u *authUsecase) Login(ctx context.Context, user domain.UserLogin) (string, error) {
	existingUser, err := u.authRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if !util.CompareHashPassword(user.Password, existingUser.Password) {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":    existingUser.Name,
			"email":   existingUser.Email,
			"expired": time.Now().Add(time.Hour * 1).Unix(),
		})

	return token.SignedString([]byte(config.C.JWTKey))
}

func (u *authUsecase) Signup(ctx context.Context, user domain.User) error {
	hashPw, err := util.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPw
	return u.authRepo.Create(ctx, user)
}
