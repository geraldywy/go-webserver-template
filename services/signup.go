package services

import (
	"errors"
	"github.com/geraldywy/go-webserver-template/database"
	"github.com/geraldywy/go-webserver-template/models"
	"github.com/go-pg/pg/v10"
)

var usernameUsedErr = errors.New("username is already taken")

func Signup(req models.SignupReq) error {
	user := &models.User{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		GroupID:        req.GroupID,
		IsAdmin: *req.IsAdmin,
	}
	_, err := database.DB.Model(user).Insert()

	if pgErr, ok := err.(pg.Error); ok {
		if pgErr.IntegrityViolation() {  // unique username constraint
			return usernameUsedErr
		}
	}
	return err
}
