package api

import (
	"encoding/json"
	"go-simple-http-server/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(s *service.StudentService) *StudentHandler {
	return &StudentHandler{service: s}
}

func (h *StudentHandler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	student, err := h.service.GetStudentById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Score int `json:"score"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if payload.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	

	student, err := h.service.CreateStudent(payload.Name, payload.Age, payload.Score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid student ID, must be an integer", http.StatusBadRequest)
        return
    }

    err = h.service.DeleteStudent(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
