package httpgin

import (
	"hse24_se_xp/app"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createUser(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createUserRequest

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := a.CreateUser(reqBody.Name, reqBody.Email, reqBody.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, UserSuccessResponse(&user))
	}
}

func createCourse(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createCourseRequest

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		course, err := a.CreateCourse(reqBody.Name, reqBody.TeacherID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, CourseSuccessResponse(&course))
	}
}

func enrollStudent(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody enrollStudentRequest

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := a.EnrollStudent(reqBody.CourseID, reqBody.StudentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student enrolled successfully"})
	}
}

func unenrollStudent(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody unenrollStudentRequest

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := a.UnenrollStudent(reqBody.CourseID, reqBody.StudentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student unenrolled successfully"})
	}
}

func createAssignment(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createAssignmentRequest

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		assignment, err := a.CreateAssignment(reqBody.CourseID, reqBody.Title, reqBody.Description, reqBody.DueDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, AssignmentSuccessResponse(&assignment))
	}
}

func submitAssignment(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		assignmentId, err := strconv.ParseInt(c.Param("assignmentId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
			return
		}

		studentId, err := strconv.ParseInt(c.Param("studentId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}

		fileData, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
			return
		}
		defer fileData.Close()

		fileBytes, err := ioutil.ReadAll(fileData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read file"})
			return
		}

		err = a.SubmitAssignment(assignmentId, studentId, fileBytes, file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Assignment submitted successfully"})
	}
}

func gradeAssignment(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody gradeAssignmentRequest

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := a.GradeAssignment(reqBody.AssignmentID, reqBody.TeacherID, reqBody.StudentID, reqBody.Grade, reqBody.Feedback)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Assignment graded successfully"})
	}
}
