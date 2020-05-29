package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dracit7/Courselect/lib/log"

	"github.com/dracit7/Courselect/lib/db"

	"github.com/dracit7/Courselect/lib/paginate"
	"github.com/dracit7/Courselect/setting"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var etime = map[string]string{
	"08:00:00": "09:40:00",
	"10:10:00": "11:50:00",
	"14:00:00": "15:40:00",
	"15:50:00": "17:30:00",
	"18:30:00": "20:10:00",
	"20:20:00": "21:50:00",
}

// StudentCourseHandler handles GET requests to /auth/student/courses.
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

// AdminCourseHandler handles GET requests to /auth/admin/applies.
func AdminCourseHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username").(string)
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	courses := db.GetAppliedCourses(page - 1)
	num := db.GetAppliedCourseNum()
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "applies.html", gin.H{
		"active":    4,
		"errors":    errors,
		"info":      info,
		"identity":  tADMIN,
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

// FacultyCourseHandler handles GET requests to /auth/faculty/courses.
func FacultyCourseHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username").(string)
	page, err := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err != nil {
		page = 0
	}

	courses := db.GetTeachingCourses(userid, page-1)
	num := db.GetTeachingCourseNum(userid)
	info := sess.Flashes("info")
	errors := sess.Flashes("error")
	sess.Save()

	c.HTML(http.StatusOK, "courses_faculty.html", gin.H{
		"active":    2,
		"errors":    errors,
		"info":      info,
		"identity":  tFACULTY,
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

// CreateCourseApply handles POST requests to /auth/faculty/courseapply.
func CreateCourseApply(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username").(string)

	credit, err := strconv.Atoi(c.PostForm("credit"))
	if err != nil {
		log.Warning(fmt.Sprintf(
			"illegal POST to server from %s", c.ClientIP()))
		return
	}

	capacity, err := strconv.Atoi(c.PostForm("capacity"))
	if err != nil {
		log.Warning(fmt.Sprintf(
			"illegal POST to server from %s", c.ClientIP()))
		return
	}

	sdate, err := strconv.Atoi(c.PostForm("sdate"))
	if err != nil {
		log.Warning(fmt.Sprintf(
			"illegal POST to server from %s", c.ClientIP()))
		return
	}

	dategap, err := strconv.Atoi(c.PostForm("dategap"))
	if err != nil {
		log.Warning(fmt.Sprintf(
			"illegal POST to server from %s", c.ClientIP()))
		return
	}

	db.CreateCourse(&db.Course{
		Name:     c.PostForm("name"),
		Teacher:  userid,
		Credit:   credit,
		Capacity: capacity,
		Sdate:    sdate,
		Edate:    sdate + dategap,
		Day:      c.PostForm("day"),
		Stime:    c.PostForm("stime"),
		Etime:    etime[c.PostForm("stime")],
		Valid:    "未通过",
	})

	sess.AddFlash("创建成功!", "info")
	sess.Save()

	c.Redirect(http.StatusFound, "/auth/faculty/courses")
	return
}

// DeleteCourseApply handles POST requests to /auth/faculty/coursedelete.
func DeleteCourseApply(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username").(string)
	courseid := c.PostForm("course")

	cid, err := strconv.Atoi(courseid)
	if err != nil {
		sess.AddFlash("删除失败: 非法的课程号", "error")
		sess.Save()
		return
	}

	db.DeleteCourse(userid, cid)
	sess.AddFlash("删除成功!", "info")
	sess.Save()

	c.Redirect(http.StatusFound, "/auth/faculty/courses")
	return
}

// PermitCourseApply handles POST requests to /auth/admin/coursepermit.
func PermitCourseApply(c *gin.Context) {
	sess := sessions.Default(c)
	courseid := c.PostForm("course")

	cid, err := strconv.Atoi(courseid)
	if err != nil {
		sess.AddFlash("审批失败: 非法的课程号", "error")
		sess.Save()
		return
	}

	db.PermitCourse(cid)
	sess.AddFlash("审批完成!", "info")
	sess.Save()

	c.Redirect(http.StatusFound, "/auth/admin/applies")
	return
}
