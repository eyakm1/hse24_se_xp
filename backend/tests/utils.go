package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hse24_se_xp/app"
	"hse24_se_xp/ports/httpgin"
	"hse24_se_xp/repo"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"time"
)

var (
	ErrBadRequest = fmt.Errorf("bad request")
	ErrForbidden  = fmt.Errorf("forbidden")
)

type userResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type courseResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	TeacherID int64  `json:"teacher_id"`
}

type usersResponse []userResponse

type assignmentResponse struct {
	ID          int64     `json:"id"`
	CourseID    int64     `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type submissionResponse struct {
	ID           int64  `json:"id"`
	AssignmentID int64  `json:"assignment_id"`
	StudentID    int64  `json:"student_id"`
	FileURL      string `json:"file_url"`
	Grade        int    `json:"grade"`
	Feedback     string `json:"feedback"`
}

type submissionsResponse []submissionResponse

type testClient struct {
	client  *http.Client
	BaseURL string
}

func GetTestClient() *testClient {
	server := httpgin.NewHTTPServer(":18080", app.NewApp(repo.New(), repo.New(), repo.New()))
	testServer := httptest.NewServer(server.Handler)

	return &testClient{
		client:  testServer.Client(),
		BaseURL: testServer.URL,
	}
}

func (tc *testClient) getResponse(req *http.Request, out any) error {
	resp, err := tc.client.Do(req)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return ErrBadRequest
		}
		if resp.StatusCode == http.StatusForbidden {
			return ErrForbidden
		}
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

func (tc *testClient) CreateUser(name, email, role string) (userResponse, error) {
	body := map[string]any{
		"name":  name,
		"email": email,
		"role":  role,
	}
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/users", tc.BaseURL), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	var resp userResponse
	err := tc.getResponse(req, &resp)
	return resp, err
}

func (tc *testClient) CreateCourse(name string, teacherID int64) (courseResponse, error) {
	body := map[string]any{
		"name":       name,
		"teacher_id": teacherID,
	}
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/courses", tc.BaseURL), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	var resp courseResponse
	err := tc.getResponse(req, &resp)
	return resp, err
}

func (tc *testClient) EnrollStudent(courseID, studentID int64) error {
	body := map[string]any{
		"course_id":  courseID,
		"student_id": studentID,
	}
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/courses/enroll", tc.BaseURL), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	return tc.getResponse(req, nil)
}

func (tc *testClient) CreateAssignment(courseID int64, title, description string, dueDate time.Time) (assignmentResponse, error) {
	body := map[string]any{
		"course_id":   courseID,
		"title":       title,
		"description": description,
		"due_date":    dueDate,
	}
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/assignments", tc.BaseURL), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	var resp assignmentResponse
	err := tc.getResponse(req, &resp)
	return resp, err
}

func (tc *testClient) SubmitAssignment(assignmentID, studentID int64, fileData []byte, fileName string) error {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", fileName)
	part.Write(fileData)
	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/assignments/%d/submit/%d", tc.BaseURL, assignmentID, studentID), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return tc.getResponse(req, nil)
}

func (tc *testClient) GradeAssignment(assignmentID, teacherID, studentID int64, grade int, feedback string) error {
	body := map[string]any{
		"assignment_id": assignmentID,
		"teacher_id":    teacherID,
		"student_id":    studentID,
		"grade":         grade,
		"feedback":      feedback,
	}
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/assignments/%d/grade", tc.BaseURL, assignmentID), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	return tc.getResponse(req, nil)
}

func (tc *testClient) ListStudents(courseID int64) (usersResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/courses/%d/students", tc.BaseURL, courseID), nil)
	req.Header.Set("Content-Type", "application/json")

	var resp usersResponse
	err := tc.getResponse(req, &resp)
	return resp, err
}

func (tc *testClient) ListSubmissions(assignmentID int64) (submissionsResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/assignments/%d/submissions", tc.BaseURL, assignmentID), nil)
	req.Header.Set("Content-Type", "application/json")

	var resp submissionsResponse
	err := tc.getResponse(req, &resp)
	return resp, err
}

func (tc *testClient) GetSubmission(assignmentID, studentID int64) (submissionResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/assignments/%d/submissions/%d", tc.BaseURL, assignmentID, studentID), nil)
	req.Header.Set("Content-Type", "application/json")

	var resp submissionResponse
	err := tc.getResponse(req, &resp)
	return resp, err
}
