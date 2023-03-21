package storge

import (
	"learning/module27/pkg/student"
	"strings"
)

type MemStorage struct {
	students map[string]student.Student
}

func NewStorage() *MemStorage {
	return &MemStorage{
		students: make(map[string]student.Student),
	}
}

func (ms *MemStorage) Get() string {
	result := "Список студетов:"
	for _, val := range ms.students {
		result = strings.Join([]string{result, val.GetInfoStr()}, "\n")
	}
	return result
}

func (ms *MemStorage) Put(name string, age, grade int) {
	ms.students[name] = student.NewStudent(name, age, grade)
}
