package httpgin

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"hse24_se_xp/app"
)

func CustomMW(c *gin.Context) {
	t := time.Now()

	c.Next()

	latency := time.Since(t)
	status := c.Writer.Status()

	log.Println("latency", latency, "method", c.Request.Method, "path", c.Request.URL.Path, "status", status)
}

func AppRouter(r *gin.RouterGroup, a app.App) {
	r.Use(CustomMW)

	// User routes
	r.POST("/users", createUser(a))
	r.PUT("/users/:user_id", updateUser(a))
	r.GET("/users/:user_id", getUser(a))
	r.DELETE("/users/:user_id", deleteUser(a))

	// Course routes
	r.POST("/courses", createCourse(a))
	r.POST("/courses/enroll", enrollStudent(a))
	r.POST("/courses/unenroll", unenrollStudent(a))
	r.GET("/teachers/:teacher_id/courses", listCourses(a))
	r.GET("/courses/:course_id/students", listStudents(a))

	// Assignment routes
	r.POST("/assignments", createAssignment(a))
	r.POST("/assignments/:assignmentId/submit/:studentId", submitAssignment(a))
	r.POST("/assignments/:assignmentId/grade", gradeAssignment(a))
	r.GET("/courses/:course_id/assignments", listAssignments(a))
	r.GET("/assignments/:assignment_id", getAssignment(a))
	r.GET("/assignments/:assignment_id/submissions", listSubmissions(a))
	r.GET("/assignments/:assignment_id/submissions/:student_id", getSubmission(a))
}
