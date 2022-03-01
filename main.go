package main

import (
	"context"
	"github.com/geraldywy/go-webserver-template/config"
	"github.com/geraldywy/go-webserver-template/database"
	"github.com/geraldywy/go-webserver-template/router"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	config.InitConfig(".", "config", "yaml")
	database.InitDB(ctx)
	defer database.CloseDB()

	r := gin.Default()
	router.InitRouting(r)

	r.Run(":8080")
}


