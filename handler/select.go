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

// SelectHandler handles GET requests to /auth/student/select.
func SelectHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username")
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	courses := db.GetSelectableCourses(page - 1)
	num := db.GetSelectableCourseNum()
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "select.html", gin.H{
		"active":    2,
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

// SelectPostHandler handles POST requests to /auth/student/select.
func SelectPostHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username")
	courseid := c.PostForm("course")

	cid, err := strconv.Atoi(courseid)
	if err != nil {
		sess.AddFlash("选课失败: 非法的课程号", "error")
		sess.Save()
		return
	}

	db.SelectCourse(userid.(string), cid)
	sess.AddFlash("选课成功!", "info")
	sess.Save()
}
