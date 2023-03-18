package student

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

func (s Student) GetInfo() (string, int, int) {
	return s.name, s.age, s.grade
}
