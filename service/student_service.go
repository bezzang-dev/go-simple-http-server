package service

import (
	"fmt"
	"go-simple-http-server/model"
)

type StudentService struct {
	db map[int]*model.Student
}

func NewStudentService() *StudentService {
	db := make(map[int]*model.Student)
	db[1] = &model.Student{Id: 1, Name: "Alice", Age: 21, Score: 90}
	db[2] = &model.Student{Id: 2, Name: "Bob", Age: 20, Score: 91}

	return &StudentService{db: db}
}

func (s *StudentService) GetStudentById(id int) (*model.Student, error) {
	student, ok := s.db[id]
	if !ok {
		return nil, fmt.Errorf("student with id %d not found", id)
	}
	return student, nil
}

func (s *StudentService) GetAllStudents() ([]*model.Student, error) {
	var students []*model.Student
	for _, student := range s.db {
		students = append(students, student)
	}
	return students, nil
}