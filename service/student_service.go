package service

import (
	"fmt"
	"go-simple-http-server/model"
)

type StudentService struct {
	db map[int]*model.Student
	nextId int
}

func NewStudentService() *StudentService {
	db := make(map[int]*model.Student)
	db[1] = &model.Student{Id: 1, Name: "Alice", Age: 21, Score: 90}
	db[2] = &model.Student{Id: 2, Name: "Bob", Age: 20, Score: 91}

	return &StudentService{db: db, nextId: 3}
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

func (s *StudentService) CreateStudent(name string, age, score int) (*model.Student, error) {

	id := s.nextId
	s.nextId++

	newStudent := &model.Student{
		Id: id,
		Name: name,
		Age: age,
		Score: score,
	}

	s.db[id] = newStudent;
	return newStudent, nil
}

func (s *StudentService) DeleteStudent(id int) error {
	if _, ok := s.db[id]; !ok {
		return fmt.Errorf("student with id %d not found", id)
	}

	delete(s.db, id)
	return nil
}