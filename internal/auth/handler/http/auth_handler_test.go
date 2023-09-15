package http

import (
	"bytes"
	"encoding/json"
	"go-template/internal/domain"
	mocks "go-template/internal/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	mockUserLogin := domain.UserLogin{
		Email:    "john@doe.com",
		Password: "123456",
	}
	mockToken := "ArPDK7pb1DgRcBy4"

	testCases := []struct {
		name         string
		expLoginCall bool
		loginErr     error
	}{
		{
			name:         "success",
			expLoginCall: true,
			loginErr:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewMockAuthUsecase(t)
			authHandler := AuthHandler{
				AuthUsecase: mockUsecase,
			}

			if tc.expLoginCall {
				mockUsecase.On("Login", mock.Anything, mock.Anything).Return(mockToken, nil)
			}

			// encode body to buffer
			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(mockUserLogin)
			assert.NoError(t, err)

			req, err := http.NewRequest("POST", "/auth/login", &body)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(authHandler.Login)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			cookie := rr.Result().Cookies()[0]
			assert.Equal(t, "token", cookie.Name)
			assert.Equal(t, mockToken, cookie.Value)
			mockUsecase.AssertExpectations(t)
		})
	}
}

func TestSignup(t *testing.T) {
	mockUser := domain.User{
		Name:     "John",
		Email:    "john@doe.com",
		Password: "123456",
	}

	testCases := []struct {
		name          string
		expSignupCall bool
		signupErr     error
	}{
		{
			name:          "success",
			expSignupCall: true,
			signupErr:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewMockAuthUsecase(t)
			authHandler := AuthHandler{
				AuthUsecase: mockUsecase,
			}

			if tc.expSignupCall {
				mockUsecase.On("Signup", mock.Anything, mock.Anything).Return(nil)
			}

			// encode body to buffer
			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(mockUser)
			assert.NoError(t, err)

			req, err := http.NewRequest("POST", "/auth/signup", &body)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(authHandler.Signup)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			mockUsecase.AssertExpectations(t)
		})
	}
}

func TestLogout(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{
			name: "success",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewMockAuthUsecase(t)
			authHandler := AuthHandler{
				AuthUsecase: mockUsecase,
			}

			req, err := http.NewRequest("POST", "/auth/logout", nil)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(authHandler.Logout)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			cookie := rr.Result().Cookies()[0]
			assert.Empty(t, cookie.Value)
			mockUsecase.AssertExpectations(t)
		})
	}
}
