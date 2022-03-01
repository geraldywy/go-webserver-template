package views

import (
	"github.com/geraldywy/go-webserver-template/models"
	"github.com/geraldywy/go-webserver-template/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginView(c *gin.Context) {
	var loginReq models.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := services.Login(loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}