package api

import (
	"encoding/json"
	"go-simple-http-server/model"
	"go-simple-http-server/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupTest() *http.ServeMux {
	studentService := service.NewStudentService()
	studentHandler := NewStudentHandler(studentService)
	router := NewRouter(studentHandler)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/", router)
	return mainRouter
}

func TestGetAllStudnets(t *testing.T) {
	router := setupTest()

	req, err := http.NewRequest(http.MethodGet, "/students", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Alice") {
		t.Errorf("handler returned unexpected body: got %v, did not contain 'Alice'", body)
	}
	if !strings.Contains(body, "Bob") {
		t.Errorf("handler returned unexpected body: got %v, did not contain 'Bob'", body)
	}

	var students []model.Student
	if err := json.Unmarshal(rr.Body.Bytes(), &students); err != nil {
		t.Fatalf("Could not unmarshal response body: %v", err)
	}
	if len(students) != 2 {
		t.Errorf("Expected 2 students, got %d", len(students))
	}
}