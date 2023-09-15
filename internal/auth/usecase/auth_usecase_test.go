package usecase

import (
	"context"
	"go-template/internal/auth/util"
	"go-template/internal/domain"
	mocks "go-template/internal/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	mockUserLogin := domain.UserLogin{
		Email:    "john@doe.com",
		Password: "123456",
	}

	hashPW, err := util.GenerateHashPassword(mockUserLogin.Password)
	assert.NoError(t, err)
	mockUser := domain.User{
		Name:     "John",
		Email:    "john@doe.com",
		Password: hashPW,
	}

	testcases := []struct {
		name              string
		expGetByEmailCall bool
		getByEmailErr     error
	}{
		{
			name:              "success",
			expGetByEmailCall: true,
			getByEmailErr:     nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewMockAuthRepository(t)

			if tc.expGetByEmailCall {
				mockRepo.On("GetByEmail", mock.Anything, mock.Anything).Return(mockUser, nil)
			}

			authUsecase := NewAuthUsecase(mockRepo)
			token, err := authUsecase.Login(context.Background(), mockUserLogin)
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		})
	}
}

func TestSignup(t *testing.T) {
	mockUser := domain.User{
		Name:     "John",
		Email:    "john@doe.com",
		Password: "123456",
	}

	testcases := []struct {
		name              string
		expGetByEmailCall bool
		getByEmailErr     error
	}{
		{
			name:              "success",
			expGetByEmailCall: true,
			getByEmailErr:     nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewMockAuthRepository(t)

			if tc.expGetByEmailCall {
				mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
			}

			authUsecase := NewAuthUsecase(mockRepo)
			err := authUsecase.Signup(context.Background(), mockUser)
			assert.NoError(t, err)
		})
	}
}
