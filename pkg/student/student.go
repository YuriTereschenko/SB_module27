package student

import "fmt"

type Student struct {
	name string

	age int

	grade int
}

func NewStudent(name string, age, grade int) Student {
	return Student{
		age:   age,
		name:  name,
		grade: grade,
	}
}

func (s Student) GetInfoStr() string {
	return fmt.Sprintf("%v %v %v", s.name, s.age, s.grade)
}
