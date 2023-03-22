package storge

import (
	"learning/module27/pkg/student"
)

type MemStorage struct {
	students map[string]student.Student
}

func NewStorage() *MemStorage {
	return &MemStorage{
		students: make(map[string]student.Student),
	}
}

func (ms *MemStorage) Get() map[string]student.Student {

	return ms.students
}

func (ms *MemStorage) Put(st student.Student) {
	ms.students[st.Name] = st
}
