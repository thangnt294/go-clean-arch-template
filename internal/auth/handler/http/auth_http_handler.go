package http

import (
	"context"
	"encoding/json"
	"go-template/internal/auth/errors"
	"go-template/internal/domain"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
	Logger      *slog.Logger
}

func NewAuthHandler(r chi.Router, us domain.AuthUsecase, logger *slog.Logger) {
	handler := &AuthHandler{
		AuthUsecase: us,
		Logger:      logger,
	}
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/signup", handler.Signup)
		r.Post("/logout", handler.Logout)
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, errors.ErrDecodeRequestBody, http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(user)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, errors.ErrValidateRequestBody, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	token, err := h.AuthUsecase.Login(ctx, user)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, errors.ErrInternalError, http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   "localhost",
		Path:     "/",
		Expires:  expirationTime,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, errors.ErrDecodeRequestBody, http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(user)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, errors.ErrValidateRequestBody, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	err = h.AuthUsecase.Signup(ctx, user)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, errors.ErrInternalError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
