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

// StudentCourseHandler handles GET requests to /auth/student/course.
func StudentCourseHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username").(string)
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	courses := db.GetSelectedCourses(userid, page-1)
	num := db.GetSelectedCourseNum(userid)
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "courses_student.html", gin.H{
		"active":    3,
		"errors":    errors,
		"info":      info,
		"identity":  tSTUDENT,
		"username":  userid,
		"courses":   courses,
		"coursenum": num,
		"start":     (page-1)*setting.UI.Pagesize + 1,
		"end":       page * setting.UI.Pagesize,
		"paginator": paginate.MakePaginator(
			c.Request.URL.Path, page, num,
		),
	})
}
