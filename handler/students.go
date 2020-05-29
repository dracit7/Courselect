package handler

import (
	"fmt"
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
	ident := sess.Get("usertype").(string)
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	students := db.GetStudents(page - 1)
	num := db.GetStudentNum()
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "students.html", gin.H{
		"active":     2,
		"errors":     errors,
		"info":       info,
		"identity":   identity[ident],
		"username":   userid,
		"students":   students,
		"studentnum": num,
		"start":      (page-1)*setting.UI.Pagesize + 1,
		"end":        page * setting.UI.Pagesize,
		"paginator": paginate.MakePaginator(
			c.Request.URL.Path+"?", page, num,
		),
	})
}

// StudentListHandler handles GET requests to /auth/studentlist.
func StudentListHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username")
	ident := sess.Get("usertype").(string)
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}
	cid, err := strconv.Atoi(c.Query("cid"))
	if err != nil {
		return
	}

	students := db.GetStudentsInCourse(cid, page-1)
	num := db.GetStudentInCourseNum(cid)
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "studentlist.html", gin.H{
		"active":     0,
		"errors":     errors,
		"info":       info,
		"identity":   identity[ident],
		"username":   userid,
		"students":   students,
		"studentnum": num,
		"start":      (page-1)*setting.UI.Pagesize + 1,
		"end":        page * setting.UI.Pagesize,
		"paginator": paginate.MakePaginator(
			fmt.Sprintf("%s?cid=%d&", c.Request.URL.Path, cid),
			page, num,
		),
	})
}
