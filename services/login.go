package services

import (
	"errors"
	"fmt"
	"github.com/geraldywy/go-webserver-template/database"
	"github.com/geraldywy/go-webserver-template/models"
)

var invalidLoginErr = errors.New("invalid login credentials")

func Login(req models.LoginReq) (*models.LoginResp, error) {
	user := new(models.User)
	if err := database.DB.Model(user).Where("username = ?", req.Username).
		Where("hashed_password = ?", req.HashedPassword).Limit(1).Select(); err != nil {
			fmt.Println(err)
			return nil, invalidLoginErr
	}

	return &models.LoginResp{
		Username: user.Username,
		GroupId:  user.GroupID,
		IsAdmin:  user.IsAdmin,
	}, nil
}