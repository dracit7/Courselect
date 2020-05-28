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
				"你需要登录才能访问本页面",
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
				fmt.Sprintf("要访问此页面，你需要是%s", ident),
				"error",
			)
			sess.Save()

			c.Redirect(http.StatusFound, "/auth/home")
			return
		}

		c.Next()
	}
}
