package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Some pages should not be accessed by tourists.
// We redirect all such requests to the login page.
func requiredLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)

		// If user's username exists in the session, we
		// regard this user as logged in. Elsewise, we
		// reject the access with a flash.
		user := sess.Get("username")
		if user == nil {
			sess.AddFlash(
				"You need to login before accessing this page.",
				"error",
			)
			sess.Save()

			c.Redirect(http.StatusFound, "/login")
			return
		}

		// If use is already logged in we allow the access.
		c.Next()
	}
}

// Identity-specific pages.
func requiredIdentity(ident string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)

		user := sess.Get("usertype")
		if user != ident {
			sess.AddFlash(
				fmt.Sprintf("You need to be %s to access this page.", ident),
				"error",
			)
			sess.Save()

			c.Redirect(http.StatusFound, "/auth/home")
			return
		}

		c.Next()
	}
}
