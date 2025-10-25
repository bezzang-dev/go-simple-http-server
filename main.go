package main

import (
	"go-simple-http-server/api"
	"go-simple-http-server/service"
	"log"
	"net/http"
)

func main() {
	studentService := service.NewStudentService()

	studentHandler := api.NewStudentHandler(studentService)

	router := api.NewRouter(studentHandler)

	log.Println("Starting server on port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}