package serve

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/info24/eva/common"
	"github.com/info24/eva/controller"
	"github.com/info24/eva/core"
	"github.com/info24/eva/middleware"
	"github.com/info24/eva/router"
	"os"
)

//go:embed static/*
var font embed.FS

func NewServer() (*gin.Engine, error) {

	port := os.Getenv(common.EvaPort)
	if len(port) == 0 {
		port = "9999"
	}

	//registerStatic()

	jwt := middleware.NewJwt()
	core.SetJwtMiddleware(jwt)

	app := gin.Default()

	router.InitRouter(app.Group("/ssh", jwt.MiddlewareFunc()))
	app.GET("/user/refresh_token", jwt.RefreshHandler)
	app.POST("/user/login", jwt.LoginHandler)
	app.GET("/user/logout", jwt.LogoutHandler)
	app.GET("/ws/:id", controller.RegisterSsh)
	app.Use(static.Serve("/", static.EmbedFolder(font, "static")))

	err := app.Run(":" + port)
	if err != nil {
		return nil, err
	}
	return app, nil
}
