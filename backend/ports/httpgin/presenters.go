package httpgin

import (
	"hse24_se_xp/app"
	"hse24_se_xp/users"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Role  users.Role `json:"role"`
}

type userResponse struct {
	ID    int64      `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Role  users.Role `json:"role"`
}

type createCourseRequest struct {
	Name      string `json:"name"`
	TeacherID int64  `json:"teacher_id"`
}

type courseResponse struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	TeacherID        int64   `json:"teacher_id"`
	EnrolledStudents []int64 `json:"enrolled_students"`
}

type enrollStudentRequest struct {
	CourseID  int64 `json:"course_id"`
	StudentID int64 `json:"student_id"`
}

type unenrollStudentRequest struct {
	CourseID  int64 `json:"course_id"`
	StudentID int64 `json:"student_id"`
}

type createAssignmentRequest struct {
	CourseID    int64     `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type assignmentResponse struct {
	ID          int64     `json:"id"`
	CourseID    int64     `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type gradeAssignmentRequest struct {
	AssignmentID int64  `json:"assignment_id"`
	TeacherID    int64  `json:"teacher_id"`
	StudentID    int64  `json:"student_id"`
	Grade        int    `json:"grade"`
	Feedback     string `json:"feedback"`
}

type submitAssignmentRequest struct {
	AssignmentID int64  `json:"assignment_id"`
	StudentID    int64  `json:"student_id"`
	FileData     []byte `json:"file_data"`
	FileName     string `json:"file_name"`
}

type submissionResponse struct {
	ID           int64  `json:"id"`
	AssignmentID int64  `json:"assignment_id"`
	StudentID    int64  `json:"student_id"`
	FileName     string `json:"file_name"`
	Grade        int    `json:"grade"`
	Feedback     string `json:"feedback"`
}

// UserSuccessResponse formats the response for a user
func UserSuccessResponse(user *users.User) *gin.H {
	return &gin.H{
		"data": userResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		"error": nil,
	}
}

// CourseSuccessResponse formats the response for a course
func CourseSuccessResponse(course *app.Course) *gin.H {
	return &gin.H{
		"data": courseResponse{
			ID:               course.ID,
			Name:             course.Name,
			TeacherID:        course.TeacherID,
			EnrolledStudents: course.EnrolledStudents,
		},
		"error": nil,
	}
}

// AssignmentSuccessResponse formats the response for an assignment
func AssignmentSuccessResponse(assignment *app.Assignment) *gin.H {
	return &gin.H{
		"data": assignmentResponse{
			ID:          assignment.ID,
			CourseID:    assignment.CourseID,
			Title:       assignment.Title,
			Description: assignment.Description,
			DueDate:     assignment.DueDate,
		},
		"error": nil,
	}
}

// SubmissionSuccessResponse formats the response for a submission
func SubmissionSuccessResponse(submission *app.Submission) *gin.H {
	return &gin.H{
		"data": submissionResponse{
			ID:           submission.ID,
			AssignmentID: submission.AssignmentID,
			StudentID:    submission.StudentID,
			FileName:     submission.FileName,
			Grade:        submission.Grade,
			Feedback:     submission.Feedback,
		},
		"error": nil,
	}
}

// UsersSuccessResponse formats the response for multiple users
func UsersSuccessResponse(users *[]users.User) *gin.H {
	var usersResponseData []userResponse
	for _, user := range *users {
		usersResponseData = append(usersResponseData, userResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return &gin.H{
		"data":  usersResponseData,
		"error": nil,
	}
}

// CoursesSuccessResponse formats the response for multiple courses
func CoursesSuccessResponse(courses *[]app.Course) *gin.H {
	var coursesResponseData []courseResponse
	for _, course := range *courses {
		coursesResponseData = append(coursesResponseData, courseResponse{
			ID:               course.ID,
			Name:             course.Name,
			TeacherID:        course.TeacherID,
			EnrolledStudents: course.EnrolledStudents,
		})
	}

	return &gin.H{
		"data":  coursesResponseData,
		"error": nil,
	}
}

// AssignmentsSuccessResponse formats the response for multiple assignments
func AssignmentsSuccessResponse(assignments *[]app.Assignment) *gin.H {
	var assignmentsResponseData []assignmentResponse
	for _, assignment := range *assignments {
		assignmentsResponseData = append(assignmentsResponseData, assignmentResponse{
			ID:          assignment.ID,
			CourseID:    assignment.CourseID,
			Title:       assignment.Title,
			Description: assignment.Description,
			DueDate:     assignment.DueDate,
		})
	}

	return &gin.H{
		"data":  assignmentsResponseData,
		"error": nil,
	}
}

// SubmissionsSuccessResponse formats the response for multiple submissions
func SubmissionsSuccessResponse(submissions *[]app.Submission) *gin.H {
	var submissionsResponseData []submissionResponse
	for _, submission := range *submissions {
		submissionsResponseData = append(submissionsResponseData, submissionResponse{
			ID:           submission.ID,
			AssignmentID: submission.AssignmentID,
			StudentID:    submission.StudentID,
			FileName:     submission.FileName,
			Grade:        submission.Grade,
			Feedback:     submission.Feedback,
		})
	}

	return &gin.H{
		"data":  submissionsResponseData,
		"error": nil,
	}
}

func UserErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
