package handler

import (
	"fmt"
	"net/http"

	"github.com/dracit7/Courselect/lib/db"
	"github.com/dracit7/Courselect/lib/log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginHandler handles GET requests to /login.
func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// LoginPostHandler handles POST requests to /login.
func LoginPostHandler(c *gin.Context) {
	var err error
	sess := sessions.Default(c)

	// Grab login information from the postform.
	username := c.PostForm("user")
	password := c.PostForm("pw")
	usertype := c.PostForm("type")

	// Compare the information provided by user and
	// the information in database.
	switch usertype {
	case "student":
		err = db.StudentLogin(username, password)
	case "faculty":
	case "admin":
	default:
		err = fmt.Errorf("identity not specified")
	}

	// Add a warning to log if someone attempted to login
	// but failed.
	if err != nil {
		log.Warning(fmt.Sprintf(
			"Failed login attempt: <%s> user %s from %s",
			usertype, username, c.ClientIP(),
		))

		// Tell frontend the type of error, and go nowhere.
		sess.AddFlash(err.Error(), "error")
		sess.Save()
		c.HTML(http.StatusOK, "login.html", gin.H{
			"errors": sess.Flashes("error"),
			"info":   sess.Flashes("info"),
		})
		sess.Clear()
		return
	}

	// If no error occurs, login succeeds.
	log.Info(fmt.Sprintf(
		"Succeeded login: <%s> %s from %s",
		usertype, username, c.ClientIP(),
	))

	// Save the username of current user to the session.
	session := sessions.Default(c)
	session.Set("username", username)
	session.Save()

	// Redirect user to the referer.
	c.Redirect(http.StatusFound, "/home")
}
