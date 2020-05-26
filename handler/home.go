package handler

import (
	"net/http"
	"strconv"

	"github.com/dracit7/Courselect/setting"

	"github.com/dracit7/Courselect/lib/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HomeHandler handles GET requests to /auth/home.
func HomeHandler(c *gin.Context) {
	sess := sessions.Default(c)
	userid := sess.Get("username")
	usertype := sess.Get("usertype")

	// Get pagination information from the request body.
	coursepage, err := strconv.Atoi(c.DefaultQuery("cp", "0"))
	if err != nil {
		coursepage = 0
	}

	// Different types of users sees different types of page.
	switch usertype {
	case "student":

		c.HTML(http.StatusOK, "home_student.html", gin.H{
			"active":    1,
			"identity":  tSTUDENT,
			"username":  db.GetStudentName(userid.(string)),
			"profile":   db.GetStudent(userid.(string)),
			"courses":   []db.Course{},
			"coursenum": 0, // TODO: Number of selected courses
		})

	case "faculty":
		courses := db.GetTeachingCourses(userid.(string), coursepage)

		c.HTML(http.StatusOK, "home_faculty.html", gin.H{
			"active":    1,
			"identity":  tFACULTY,
			"username":  db.GetFacultyName(userid.(string)),
			"profile":   db.GetFaculty(userid.(string)),
			"courses":   courses,
			"coursenum": len(courses),
		})

	case "admin":

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"active":   1,
			"identity": tADMIN,
			"username": setting.Admin.Username,
		})

	}
}
