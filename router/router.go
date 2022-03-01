package router

import (
	"github.com/geraldywy/go-webserver-template/views"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine) {
	r.GET("/ping", views.Ping)
	r.POST("/api/signup", views.SignupView)
	r.POST("/api/login", views.LoginView)
}