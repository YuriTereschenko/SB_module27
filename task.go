package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string

	age int

	grade int
}

func newStudent(name string, age, grade int) Student {
	return Student{
		age:   age,
		name:  name,
		grade: grade,
	}
}

func (s Student) getInfo() {
	fmt.Println(s.name, s.age, s.grade)
}

func get(storage map[string]Student) {
	fmt.Println("Студенты из хранилища:")
	for _, value := range storage {
		value.getInfo()
	}
}

func put(data string, storage map[string]Student) {
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

	storage[name] = newStudent(name, age, grade)

}

func main() {
	studStorage := make(map[string]Student)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		put(txt, studStorage)
	}
	get(studStorage)
}
