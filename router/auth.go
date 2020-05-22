package router

import (
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
