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

// TestCreateStudent tests the POST /students endpoint
func TestCreateStudent(t *testing.T) {
	router := setupTest()

	tests := []struct {
		name       string
		payload    string
		wantStatus int
		wantName   string
	}{
		{
			name:       "Happy Path - Create Student",
			payload:    `{"name":"Charlie","age":23, "score":90}`,
			wantStatus: http.StatusCreated,
			wantName:   "Charlie",
		},
		{
			name:       "Error Path - Invalid JSON",
			payload:    `{"name":"David"`, // Malformed JSON
			wantStatus: http.StatusBadRequest,
			wantName:   "",
		},
		{
			name:       "Error Path - Missing Name",
			payload:    `{"age":25}`, // Name is required
			wantStatus: http.StatusBadRequest,
			wantName:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/students", strings.NewReader(tc.payload))
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if rr.Code != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tc.wantStatus)
			}

			if tc.wantStatus == http.StatusCreated {
				var student model.Student
				if err := json.Unmarshal(rr.Body.Bytes(), &student); err != nil {
					t.Fatalf("Could not unmarshal response: %v", err)
				}
				if student.Name != tc.wantName {
					t.Errorf("handler returned unexpected name: got %v want %v",
						student.Name, tc.wantName)
				}

				// Check if the ID is "3" (since we have 2 mock students)
				if student.Id != 3 {
					t.Errorf("handler returned unexpected ID: got %v want 3", student.Id)
				}
			}
		})
	}
}

// TestDeleteStudent tests the DELETE /students/{id} endpoint
func TestDeleteStudent(t *testing.T) {
	
	router := setupTest()

	t.Run("Happy Path - Delete Student", func(t *testing.T) {
		reqDelete, _ := http.NewRequest(http.MethodDelete, "/students/1", nil)
		rrDelete := httptest.NewRecorder()

		router.ServeHTTP(rrDelete, reqDelete)

		if rrDelete.Code != http.StatusNoContent {
			t.Errorf("DELETE handler returned wrong status code: got %v want %v",
				rrDelete.Code, http.StatusNoContent)
		}

		reqGet, _ := http.NewRequest(http.MethodGet, "/students/1", nil)
		rrGet := httptest.NewRecorder()

		// Serve the GET request *on the same router*
		router.ServeHTTP(rrGet, reqGet)

		// Check that we now get a 404 Not Found
		if rrGet.Code != http.StatusNotFound {
			t.Errorf("GET handler after delete returned wrong status: got %v want %v",
				rrGet.Code, http.StatusNotFound)
		}
	})

	t.Run("Error Path - Delete Non-existent Student", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/students/999", nil)
		rr := httptest.NewRecorder()

		routerFresh := setupTest()
		routerFresh.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("DELETE handler for non-existent student returned wrong status: got %v want %v",
				rr.Code, http.StatusNotFound)
		}
	})
}
