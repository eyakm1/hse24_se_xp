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

func NewApp(userRepo, courseRepo, submissionRepo Repository) App {
	return &HomeworkService{
		users:       userRepo,
		courses:     courseRepo,
		submissions: submissionRepo,
	}
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
	users       Repository
	courses     Repository
	submissions Repository
}

var PermissionDenied = errors.New("the user does not have enough permission to perform this action")
var DefunctUser = errors.New("there is no user with this ID")

func (h *HomeworkService) CreateUser(name string, email string, role users.Role) (users.User, error) {
	user := users.User{ID: h.users.GetNextId(), Name: name, Email: email, Role: role}

	return user, h.users.Add(user)
}

func (h *HomeworkService) UpdateUser(userId int64, name string, email string) (users.User, error) {
	if !h.users.CheckIdExist(userId) {
		return users.User{}, DefunctUser
	}

	res, err := h.users.Get(userId)
	if err != nil {
		return users.User{}, err
	}

	user := res.(users.User)
	user.Name = name
	user.Email = email

	return user, h.users.Update(userId, user)
}

func (h *HomeworkService) GetUser(userId int64) (users.User, error) {
	if !h.users.CheckIdExist(userId) {
		return users.User{}, DefunctUser
	}

	res, err := h.users.Get(userId)
	if err != nil {
		return users.User{}, err
	}

	return res.(users.User), nil
}

func (h *HomeworkService) DeleteUser(userId int64) error {
	if !h.users.CheckIdExist(userId) {
		return DefunctUser
	}

	return h.users.Delete(userId)
}

func (h *HomeworkService) CreateCourse(name string, teacherId int64) (Course, error) {
	course := Course{ID: h.courses.GetNextId(), Name: name, TeacherID: teacherId}
	return course, h.courses.Add(course)
}

func (h *HomeworkService) EnrollStudent(courseId int64, studentId int64) error {
	if !h.courses.CheckIdExist(courseId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}

	res, err := h.courses.Get(courseId)
	if err != nil {
		return err
	}

	course := res.(Course)
	course.EnrolledStudents = append(course.EnrolledStudents, studentId)

	return h.courses.Update(courseId, course)
}

func (h *HomeworkService) UnenrollStudent(courseId int64, studentId int64) error {
	if !h.courses.CheckIdExist(courseId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}

	res, err := h.courses.Get(courseId)
	if err != nil {
		return err
	}

	course := res.(Course)
	for i, id := range course.EnrolledStudents {
		if id == studentId {
			course.EnrolledStudents = append(course.EnrolledStudents[:i], course.EnrolledStudents[i+1:]...)
			break
		}
	}

	return h.courses.Update(courseId, course)
}

func (h *HomeworkService) ListCourses(teacherId int64) ([]Course, error) {
	if !h.users.CheckIdExist(teacherId) {
		return nil, DefunctUser
	}

	var courses []Course
	for _, item := range h.courses.GetArray() {
		course := item.(Course)
		if course.TeacherID == teacherId {
			courses = append(courses, course)
		}
	}
	return courses, nil
}

func (h *HomeworkService) ListStudents(courseId int64) ([]users.User, error) {
	if !h.courses.CheckIdExist(courseId) {
		return nil, DefunctUser
	}

	res, err := h.courses.Get(courseId)
	if err != nil {
		return nil, err
	}

	course := res.(Course)
	var students []users.User
	for _, studentId := range course.EnrolledStudents {
		studentRes, err := h.users.Get(studentId)
		if err != nil {
			return nil, err
		}
		students = append(students, studentRes.(users.User))
	}
	return students, nil
}

func (h *HomeworkService) CreateAssignment(courseId int64, title string, description string, dueDate time.Time) (Assignment, error) {
	if !h.courses.CheckIdExist(courseId) {
		return Assignment{}, DefunctUser
	}

	assignment := Assignment{
		ID:          h.courses.GetNextId(),
		CourseID:    courseId,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}

	err := h.courses.Add(assignment)
	if err != nil {
		return Assignment{}, err
	}

	return assignment, nil
}

func (h *HomeworkService) SubmitAssignment(assignmentId int64, studentId int64, fileData []byte, fileName string) error {
	if !h.courses.CheckIdExist(assignmentId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}

	submission := Submission{
		ID:           h.submissions.GetNextId(),
		AssignmentID: assignmentId,
		StudentID:    studentId,
		FileData:     fileData,
		FileName:     fileName,
	}

	return h.submissions.Add(submission)
}

func (h *HomeworkService) GradeAssignment(assignmentId int64, teacherId int64, studentId int64, grade int, feedback string) error {
	if !h.courses.CheckIdExist(assignmentId) || !h.users.CheckIdExist(studentId) {
		return DefunctUser
	}

	res, err := h.submissions.Get(assignmentId)
	if err != nil {
		return err
	}

	submission := res.(Submission)
	submission.Grade = grade
	submission.Feedback = feedback

	return h.submissions.Update(assignmentId, submission)
}

func (h *HomeworkService) ListAssignments(courseId int64) ([]Assignment, error) {
	if !h.courses.CheckIdExist(courseId) {
		return nil, DefunctUser
	}

	var assignments []Assignment
	for _, item := range h.courses.GetArray() {
		assignment := item.(Assignment)
		if assignment.CourseID == courseId {
			assignments = append(assignments, assignment)
		}
	}
	return assignments, nil
}

func (h *HomeworkService) GetAssignment(assignmentId int64) (Assignment, error) {
	if !h.courses.CheckIdExist(assignmentId) {
		return Assignment{}, DefunctUser
	}

	res, err := h.courses.Get(assignmentId)
	if err != nil {
		return Assignment{}, err
	}

	return res.(Assignment), nil
}

func (h *HomeworkService) ListSubmissions(assignmentId int64) ([]Submission, error) {
	if !h.courses.CheckIdExist(assignmentId) {
		return nil, DefunctUser
	}

	var submissions []Submission
	for _, item := range h.submissions.GetArray() {
		submission := item.(Submission)
		if submission.AssignmentID == assignmentId {
			submissions = append(submissions, submission)
		}
	}
	return submissions, nil
}

func (h *HomeworkService) GetSubmission(assignmentId int64, studentId int64) (Submission, error) {
	if !h.courses.CheckIdExist(assignmentId) || !h.users.CheckIdExist(studentId) {
		return Submission{}, DefunctUser
	}

	res, err := h.submissions.Get(assignmentId)
	if err != nil {
		return Submission{}, err
	}

	return res.(Submission), nil
}
