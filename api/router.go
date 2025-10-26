package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handler *StudentHandler) *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/students", handler.GetAllStudents).Methods(http.MethodGet)
	mux.HandleFunc("/students/{id}", handler.GetStudentById).Methods(http.MethodGet)

	mux.HandleFunc("/students", handler.CreateStudent).Methods(http.MethodPost)
	mux.HandleFunc("/students/{id}", handler.DeleteStudent).Methods(http.MethodDelete)

	return mux
}