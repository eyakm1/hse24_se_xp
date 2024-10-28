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
	assert.Equal(t, "Test User", createdUser.Name)
	assert.Equal(t, "test@testing.ru", createdUser.Email)
	assert.Equal(t, "student", createdUser.Role)
}

func TestCreateCourse(t *testing.T) {
	client := GetTestClient()

	createdTeacher, err := client.CreateUser("Test Teacher", "teacher@testing.ru", "teacher")
	assert.NoError(t, err)

	course, err := client.CreateCourse("Test Course", createdTeacher.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Test Course", course.Name)
	assert.Equal(t, createdTeacher.ID, course.TeacherID)
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
	assert.Contains(t, students.Data, createdStudent)
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
	assert.Equal(t, "Test Assignment", assignment.Title)
	assert.Equal(t, "This is a test assignment", assignment.Description)
	assert.Equal(t, course.ID, assignment.CourseID)
	assert.WithinDuration(t, dueDate, assignment.DueDate, time.Second)
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
	assert.NotEmpty(t, submissions.Data)
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
	assert.Equal(t, 95, submission.Grade)
	assert.Equal(t, "Great job!", submission.Feedback)
}
