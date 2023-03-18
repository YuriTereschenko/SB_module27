package storge

import (
	"fmt"
	"learning/module27/pkg/student"
	"strconv"
	"strings"
)

func NewStorage() map[string]student.Student {
	return make(map[string]student.Student)
}

func Get(storage map[string]student.Student) {
	fmt.Println("Студенты из хранилища:")
	for _, value := range storage {
		fmt.Println(value.GetInfo())
	}
}

func Put(data string, storage map[string]student.Student) {
	dataSplited := strings.Split(data, " ")
	if len(dataSplited) != 3 {
		fmt.Println("Incorrect input")
		return
	}
	name := dataSplited[0]
	age, err := strconv.Atoi(dataSplited[1])
	if err != nil {
		fmt.Println("Incorrect input", err)
		return
	}
	grade, err := strconv.Atoi(dataSplited[2])
	if err != nil {
		fmt.Println("Incorrect input", err)
		return
	}

	storage[name] = student.NewStudent(name, age, grade)

}
