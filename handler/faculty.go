package handler

import (
	"net/http"
	"strconv"

	"github.com/dracit7/Courselect/lib/db"
	"github.com/dracit7/Courselect/lib/paginate"
	"github.com/dracit7/Courselect/setting"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// FacultyHandler handles GET requests to /auth/admin/faculty.
func FacultyHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username")
	ident := sess.Get("usertype").(string)
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	faculty := db.GetFacultys(page - 1)
	num := db.GetFacultyNum()
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "faculty.html", gin.H{
		"active":     3,
		"errors":     errors,
		"info":       info,
		"identity":   identity[ident],
		"username":   userid,
		"faculty":    faculty,
		"facultynum": num,
		"start":      (page-1)*setting.UI.Pagesize + 1,
		"end":        page * setting.UI.Pagesize,
		"paginator": paginate.MakePaginator(
			c.Request.URL.Path+"?", page, num,
		),
	})
}
