package views

import (
	"github.com/geraldywy/go-webserver-template/models"
	"github.com/geraldywy/go-webserver-template/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignupView(c *gin.Context) {
	var signupReq models.SignupReq
	if err := c.ShouldBindJSON(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.Signup(signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"username": signupReq.Username, "group_id": signupReq.GroupID, "is_admin": signupReq.IsAdmin})
}