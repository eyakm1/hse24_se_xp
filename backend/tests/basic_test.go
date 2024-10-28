package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	client := GetTestClient()

	createdUser, err := client.CreateUser("Test User", "test@testing.ru", "student")
	assert.NoError(t, err)
	assert.Equal(t, createdUser.Name, "Test User")
	assert.Equal(t, createdUser.Email, "test@testing.ru")
	assert.Equal(t, createdUser.Role, "student")
}

func TestCreateCourse(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", "teacher")
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.ID)
	assert.NoError(t, err)
	assert.Equal(t, course.Name, "Test Course")
	assert.Equal(t, course.TeacherID, createdTeacher.ID)
}

func TestEnrollStudent(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", "teacher")
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.ID)
	assert.NoError(t, err)

	createdStudent, err := client.CreateUser("Test Student", "student@testing.ru", "student")
	assert.NoError(t, err)

	err = client.EnrollStudent(course.ID, createdStudent.ID)
	assert.NoError(t, err)

	students, err := client.ListStudents(course.ID)
	assert.NoError(t, err)
	assert.Contains(t, students, createdStudent)
}

func TestCreateAssignment(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", "teacher")
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.ID)
	assert.NoError(t, err)

	dueDate := time.Now().AddDate(0, 0, 7)
	assignment, err := client.CreateAssignment(course.ID, "Test Assignment", "This is a test assignment", dueDate)
	assert.NoError(t, err)
	assert.Equal(t, assignment.Title, "Test Assignment")
	assert.Equal(t, assignment.Description, "This is a test assignment")
	assert.Equal(t, assignment.CourseID, course.ID)
	assert.WithinDuration(t, assignment.DueDate, dueDate, time.Second)
}

func TestSubmitAssignment(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", "teacher")
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.ID)
	assert.NoError(t, err)

	createdStudent, err := client.CreateUser("Test Student", "student@testing.ru", "student")
	assert.NoError(t, err)

	err = client.EnrollStudent(course.ID, createdStudent.ID)
	assert.NoError(t, err)

	dueDate := time.Now().AddDate(0, 0, 7)
	assignment, err := client.CreateAssignment(course.ID, "Test Assignment", "This is a test assignment", dueDate)
	assert.NoError(t, err)

	fileData := []byte("This is the content of the assignment.")
	fileName := "assignment.pdf"
	err = client.SubmitAssignment(assignment.ID, createdStudent.ID, fileData, fileName)
	assert.NoError(t, err)

	submissions, err := client.ListSubmissions(assignment.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, submissions)
}

func TestGradeAssignment(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", "teacher")
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.ID)
	assert.NoError(t, err)

	createdStudent, err := client.CreateUser("Test Student", "student@testing.ru", "student")
	assert.NoError(t, err)

	err = client.EnrollStudent(course.ID, createdStudent.ID)
	assert.NoError(t, err)

	dueDate := time.Now().AddDate(0, 0, 7)
	assignment, err := client.CreateAssignment(course.ID, "Test Assignment", "This is a test assignment", dueDate)
	assert.NoError(t, err)

	fileData := []byte("This is the content of the assignment.")
	fileName := "assignment.pdf"
	err = client.SubmitAssignment(assignment.ID, createdStudent.ID, fileData, fileName)
	assert.NoError(t, err)

	err = client.GradeAssignment(assignment.ID, createdTeacher.ID, createdStudent.ID, 95, "Great job!")
	assert.NoError(t, err)

	submission, err := client.GetSubmission(assignment.ID, createdStudent.ID)
	assert.NoError(t, err)
	assert.Equal(t, submission.Grade, 95)
	assert.Equal(t, submission.Feedback, "Great job!")
}
