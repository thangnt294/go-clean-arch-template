package mysql

import (
	"context"
	"fmt"
	"go-template/config"
	"go-template/internal/domain"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetByEmail(t *testing.T) {
	config.LoadConfig("./../../../../.env")
	fmt.Println(config.C)
	db := sqlx.MustConnect(config.C.DBDriver, config.C.DBUrl)
	authRepo := NewAuthRepo(db)

	user := domain.User{
		Name:     "John",
		Email:    "john@doe.com",
		Password: "123456",
	}
	err := authRepo.Create(context.Background(), user)
	assert.NoError(t, err)
	existingUser, err := authRepo.GetByEmail(context.Background(), user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, existingUser.Name)
	assert.Equal(t, user.Email, existingUser.Email)
}

func TestCreate(t *testing.T) {
	config.LoadConfig("./../../../../.env")
	fmt.Println(config.C)
	db := sqlx.MustConnect(config.C.DBDriver, config.C.DBUrl)
	authRepo := NewAuthRepo(db)

	user := domain.User{
		Name:     "John",
		Email:    "john@doe.com",
		Password: "123456",
	}
	err := authRepo.Create(context.Background(), user)
	assert.NoError(t, err)
}
