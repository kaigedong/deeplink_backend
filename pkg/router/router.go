package router

import (
	"github.com/gin-gonic/gin"

	"deeplink_backend/pkg/ctrl"
	"deeplink_backend/pkg/log"
	"deeplink_backend/pkg/middleware"
)

func ConfigRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	// Gin framework will also output it's log
	gin.DefaultWriter = log.Log.Out

	router := gin.Default()
	router.GET("/404", ctrl.Engines.NotFound)
	router.GET("/ok", ctrl.Engines.OK)
	// engine.POST(, handlers ...gin.HandlerFunc)

	authMiddleware, err := middleware.JwtMiddleware()
	if err != nil {
		// TODO: add log info
		// log.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		router.PUT("/new_device", ctrl.Engines.NewDevice)
		router.DELETE("/delete_device", ctrl.Engines.NewDevice)
		router.GET("/device", ctrl.Engines.Devides)
	}

	return router
}
