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

// StudentHandler handles GET requests to /auth/admin/students.
func StudentHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username")
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	students := db.GetStudents(page - 1)
	num := db.GetStudentNum()
	c.HTML(http.StatusOK, "students.html", gin.H{
		"active":     2,
		"identity":   tADMIN,
		"username":   userid,
		"students":   students,
		"studentnum": num,
		"start":      (page-1)*setting.UI.Pagesize + 1,
		"end":        page * setting.UI.Pagesize,
		"paginator": paginate.MakePaginator(
			c.Request.URL.Path, page, num,
		),
	})
}
