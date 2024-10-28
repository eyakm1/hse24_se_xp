package tests

import (
	"hse24_se_xp/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	client := GetTestClient()

	createdUser, err := client.CreateUser("Test User", "test@testing.ru", 0)
	assert.NoError(t, err)
	assert.Equal(t, "Test User", createdUser.Data.Name)
	assert.Equal(t, "test@testing.ru", createdUser.Data.Email)
	assert.Equal(t, users.Role(0), createdUser.Data.Role)
}

func TestCreateCourse(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", 1)
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.Data.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Test Course", course.Data.Name)
	assert.Equal(t, createdTeacher.Data.ID, course.Data.TeacherID)
}

func TestEnrollStudent(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", 1)
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.Data.ID)
	assert.NoError(t, err)

	createdStudent, err := client.CreateUser("Test Student", "student@testing.ru", 0)
	assert.NoError(t, err)

	err = client.EnrollStudent(course.Data.ID, createdStudent.Data.ID)
	assert.NoError(t, err)

	students, err := client.ListStudents(course.Data.ID)
	assert.NoError(t, err)
	assert.Contains(t, students.Data, createdStudent)
}

func TestCreateAssignment(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", 1)
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.Data.ID)
	assert.NoError(t, err)

	dueDate := time.Now().AddDate(0, 0, 7)
	assignment, err := client.CreateAssignment(course.Data.ID, "Test Assignment", "This is a test assignment", dueDate)
	assert.NoError(t, err)
	assert.Equal(t, "Test Assignment", assignment.Data.Title)
	assert.Equal(t, "This is a test assignment", assignment.Data.Description)
	assert.Equal(t, course.Data.ID, assignment.Data.CourseID)
	assert.WithinDuration(t, dueDate, assignment.Data.DueDate, time.Second)
}

// func TestSubmitAssignment(t *testing.T) {
// 	client := GetTestClient()

// 	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", 1)
// 	assert.NoError(t, err)

// 	course, err := client.CreateCourse("Test Course", createdTeacher.Data.ID)
// 	assert.NoError(t, err)

// 	createdStudent, err := client.CreateUser("Test Student", "student@testing.ru", 0)
// 	assert.NoError(t, err)

// 	err = client.EnrollStudent(course.Data.ID, createdStudent.Data.ID)
// 	assert.NoError(t, err)

// 	dueDate := time.Now().AddDate(0, 0, 7)
// 	assignment, err := client.CreateAssignment(course.Data.ID, "Test Assignment", "This is a test assignment", dueDate)
// 	assert.NoError(t, err)

// 	fileData := []byte("This is the content of the assignment.")
// 	fileName := "assignment.pdf"
// 	err = client.SubmitAssignment(assignment.Data.ID, createdStudent.Data.ID, fileData, fileName)
// 	assert.NoError(t, err)

// 	submissions, err := client.ListSubmissions(assignment.Data.ID)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, submissions.Data)
// }

// func TestGradeAssignment(t *testing.T) {
// 	client := GetTestClient()

// 	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", 1)
// 	assert.NoError(t, err)

// 	course, err := client.CreateCourse("Test Course", createdTeacher.Data.ID)
// 	assert.NoError(t, err)

// 	createdStudent, err := client.CreateUser("Test Student", "student@testing.ru", 0)
// 	assert.NoError(t, err)

// 	err = client.EnrollStudent(course.Data.ID, createdStudent.Data.ID)
// 	assert.NoError(t, err)

// 	dueDate := time.Now().AddDate(0, 0, 7)
// 	assignment, err := client.CreateAssignment(course.Data.ID, "Test Assignment", "This is a test assignment", dueDate)
// 	assert.NoError(t, err)

// 	fileData := []byte("This is the content of the assignment.")
// 	fileName := "assignment.pdf"
// 	err = client.SubmitAssignment(assignment.Data.ID, createdStudent.Data.ID, fileData, fileName)
// 	assert.NoError(t, err)

// 	err = client.GradeAssignment(assignment.Data.ID, createdTeacher.Data.ID, createdStudent.Data.ID, 95, "Great job!")
// 	assert.NoError(t, err)

// 	submission, err := client.GetSubmission(assignment.Data.ID, createdStudent.Data.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 95, submission.Data.Grade)
// 	assert.Equal(t, "Great job!", submission.Data.Feedback)
// }
