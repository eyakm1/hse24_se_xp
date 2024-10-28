package app

import (
	"hse24_se_xp/users"
	"time"

	"github.com/pkg/errors"
)

type App interface {
	CreateUser(name string, email string, role users.Role) (users.User, error)
	UpdateUser(userId int64, name string, email string) (users.User, error)
	GetUser(userId int64) (users.User, error)
	DeleteUser(userId int64) error

	// Course methods
	CreateCourse(name string, teacherId int64) (Course, error)
	EnrollStudent(courseId int64, studentId int64) error
	UnenrollStudent(courseId int64, studentId int64) error
	ListCourses(teacherId int64) ([]Course, error)
	ListStudents(courseId int64) ([]users.User, error)

	// Assignment methods
	CreateAssignment(courseId int64, title string, description string, dueDate time.Time) (Assignment, error)
	SubmitAssignment(assignmentId int64, studentId int64, fileData []byte, fileName string) error
	GradeAssignment(assignmentId int64, teacherId int64, studentId int64, grade int, feedback string) error
	ListAssignments(courseId int64) ([]Assignment, error)
	GetAssignment(assignmentId int64) (Assignment, error)
	ListSubmissions(assignmentId int64) ([]Submission, error)
	GetSubmission(assignmentId int64, studentId int64) (Submission, error)
}

func NewApp(userRepo Repository) App {
	return &HomeworkService{users: userRepo}
}

type Repository interface {
	Add(e interface{}) error
	Update(id int64, e interface{}) error
	Get(id int64) (interface{}, error)
	Delete(id int64) error
	CheckIdExist(id int64) bool
	GetNextId() int64
	GetArray() []interface{}
}

type Course struct {
	ID               int64
	Name             string
	TeacherID        int64
	EnrolledStudents []int64
}

type Assignment struct {
	ID          int64
	CourseID    int64
	Title       string
	Description string
	DueDate     time.Time
}

type Submission struct {
	ID           int64
	AssignmentID int64
	StudentID    int64
	FileData     []byte
	FileName     string
	Grade        int
	Feedback     string
}

type HomeworkService struct {
	users Repository
}

var PermissionDenied = errors.New("the user does not have enough permission to edit the ad")
var DefunctUser = errors.New("there is no user with this ID")

func (h *HomeworkService) CreateUser(name string, email string, role users.Role) (users.User, error) {
	user := users.User{ID: h.users.GetNextId(), Name: name, Email: email, Role: role}

	return user, h.users.Add(user)
}

func (u *HomeworkService) UpdateUser(userId int64, name string, email string) (users.User, error) {
	if !u.users.CheckIdExist(userId) {
		return users.User{}, DefunctUser
	}

	res, err := u.users.Get(userId)
	user := res.(users.User)

	if err != nil {
		return user, err
	}

	user.Name = name
	user.Email = email

	return user, u.users.Update(userId, user)
}

func (u *HomeworkService) GetUser(userId int64) (users.User, error) {
	if !u.users.CheckIdExist(userId) {
		return users.User{}, DefunctUser
	}

	res, err := u.users.Get(userId)
	if err != nil {
		return users.User{}, err
	}

	return res.(users.User), nil
}

func (u *HomeworkService) DeleteUser(userId int64) error {
	if !u.users.CheckIdExist(userId) {
		return DefunctUser
	}

	return u.users.Delete(userId)
}

func (h *HomeworkService) CreateCourse(name string, teacherId int64) (Course, error) {
	course := Course{ID: h.users.GetNextId(), Name: name, TeacherID: teacherId}
	return course, h.users.Add(course)
}

func (h *HomeworkService) EnrollStudent(courseId int64, studentId int64) error {
	if !h.users.CheckIdExist(courseId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}
	// Assuming we have a method to add student to course
	return h.users.Update(courseId, studentId)
}

func (h *HomeworkService) CreateAssignment(courseId int64, title string, description string, dueDate time.Time) (Assignment, error) {
	if !h.users.CheckIdExist(courseId) {
		return Assignment{}, DefunctUser
	}

	assignment := Assignment{
		ID:          h.users.GetNextId(),
		CourseID:    courseId,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}

	err := h.users.Add(assignment)
	if err != nil {
		return Assignment{}, err
	}

	return assignment, nil
}

// SubmitAssignment implements App.
func (h *HomeworkService) SubmitAssignment(assignmentId int64, studentId int64, fileData []byte, fileName string) error {
	if !h.users.CheckIdExist(assignmentId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}

	submission := Submission{
		ID:           h.users.GetNextId(),
		AssignmentID: assignmentId,
		StudentID:    studentId,
		FileData:     fileData,
		FileName:     fileName,
	}

	return h.users.Add(submission)
}

func (h *HomeworkService) GetAssignment(assignmentId int64) (Assignment, error) {
	if !h.users.CheckIdExist(assignmentId) {
		return Assignment{}, DefunctUser
	}
	res, err := h.users.Get(assignmentId)
	if err != nil {
		return Assignment{}, err
	}
	return res.(Assignment), nil
}

func (h *HomeworkService) GetSubmission(assignmentId int64, studentId int64) (Submission, error) {
	if !h.users.CheckIdExist(assignmentId) || !h.users.CheckIdExist(studentId) {
		return Submission{}, DefunctUser
	}
	res, err := h.users.Get(assignmentId)
	if err != nil {
		return Submission{}, err
	}
	return res.(Submission), nil
}

func (h *HomeworkService) GradeAssignment(assignmentId int64, teacherId int64, studentId int64, grade int, feedback string) error {
	if !h.users.CheckIdExist(assignmentId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}
	res, err := h.users.Get(assignmentId)
	if err != nil {
		return err
	}
	submission := res.(Submission)
	submission.Grade = grade
	submission.Feedback = feedback
	return h.users.Update(assignmentId, submission)
}

func (h *HomeworkService) ListAssignments(courseId int64) ([]Assignment, error) {
	if !h.users.CheckIdExist(courseId) {
		return nil, DefunctUser
	}
	var assignments []Assignment
	for _, item := range h.users.GetArray() {
		assignment := item.(Assignment)
		if assignment.CourseID == courseId {
			assignments = append(assignments, assignment)
		}
	}
	return assignments, nil
}

func (h *HomeworkService) ListCourses(teacherId int64) ([]Course, error) {
	if !h.users.CheckIdExist(teacherId) {
		return nil, DefunctUser
	}
	var courses []Course
	for _, item := range h.users.GetArray() {
		course := item.(Course)
		if course.TeacherID == teacherId {
			courses = append(courses, course)
		}
	}
	return courses, nil
}

func (h *HomeworkService) ListStudents(courseId int64) ([]users.User, error) {
	if !h.users.CheckIdExist(courseId) {
		return nil, DefunctUser
	}
	var students []users.User
	for _, item := range h.users.GetArray() {
		student := item.(users.User)
		if student.Role == users.Student {
			students = append(students, student)
		}
	}
	return students, nil
}

func (h *HomeworkService) ListSubmissions(assignmentId int64) ([]Submission, error) {
	if !h.users.CheckIdExist(assignmentId) {
		return nil, DefunctUser
	}
	var submissions []Submission
	for _, item := range h.users.GetArray() {
		submission := item.(Submission)
		if submission.AssignmentID == assignmentId {
			submissions = append(submissions, submission)
		}
	}
	return submissions, nil
}

// UnenrollStudent implements App.
func (h *HomeworkService) UnenrollStudent(courseId int64, studentId int64) error {
	if !h.users.CheckIdExist(courseId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}

	// Assuming we have a method to remove student from course
	// This is a placeholder implementation and should be replaced with actual logic
	course, err := h.users.Get(courseId)
	if err != nil {
		return err
	}

	// Assuming course has a list of enrolled students
	courseData := course.(Course)
	for i, id := range courseData.EnrolledStudents {
		if id == studentId {
			courseData.EnrolledStudents = append(courseData.EnrolledStudents[:i], courseData.EnrolledStudents[i+1:]...)
			break
		}
	}

	return h.users.Update(courseId, courseData)
}
