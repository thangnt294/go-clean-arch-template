package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-template/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(r chi.Router, us domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: us,
	}
	r.Post("/login", handler.Login)
	r.Post("/signup", handler.Signup)
	r.Post("/logout", handler.Logout)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err) // TODO: add logging
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := h.AuthUsecase.Login(context.Background(), user)
	if err != nil {
		fmt.Println(err) // TODO: add logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   60 * 60, // TODO: set this to expiration date
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
		fmt.Println(err) // TODO: add logging
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.AuthUsecase.Signup(context.Background(), user)
	if err != nil {
		fmt.Println(err) // TODO: add logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
