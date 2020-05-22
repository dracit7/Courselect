package router

import (
	"github.com/dracit7/Courselect/handler"
	"github.com/dracit7/Courselect/setting"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Setup sets up the router of gin server.
//
// This router dispatches all requests from clients
// to registered handlers with registered middlewares.
func Setup() *gin.Engine {
	gin.SetMode(setting.Server.Mode)

	// Create and configure the router.
	router := gin.New()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/**/*")

	// Use cookies to store current session.
	cookies := sessions.NewCookieStore(
		[]byte(uuid.New().String()),
	)

	// Set some global middlewares.
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions(uuid.New().String(), cookies))

	// All pages should be logged in to access except
	// the login page and the home page.
	auth := router.Group("/auth")
	auth.Use(requiredLoggedIn())

	auth.GET("/home", handler.HomeHandler)

	// Services requires login too.
	service := router.Group("/services")
	service.Use(requiredLoggedIn())

	service.GET("/logout", handler.LogoutHandler)

	router.GET("/login", handler.LoginHandler)
	router.POST("/login", handler.LoginPostHandler)

	return router
}
