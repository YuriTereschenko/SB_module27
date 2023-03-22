package main

import (
	"bufio"
	"fmt"
	"learning/module27/pkg/storge"
	"learning/module27/pkg/student"
	"os"
	"strconv"
	"strings"
)

type Storage interface {
	Put(string, int, int)
	Get() string
}

func convertData(data string) (bool, student.Student) {
	dataSplited := strings.Split(data, " ")
	if len(dataSplited) != 3 {
		fmt.Println("Incorrect input")
		return false, student.Student{}
	}
	name := dataSplited[0]
	age, err := strconv.Atoi(dataSplited[1])
	if err != nil {
		fmt.Println("Incorrect input", err)
		return false, student.Student{}
	}
	grade, err := strconv.Atoi(dataSplited[2])
	if err != nil {
		fmt.Println("Incorrect input", err)
		return false, student.Student{}
	}
	return true, student.Student{name, age, grade}
}
func printStudData(data map[string]student.Student) {
	result := "Список студетов:"
	for _, val := range data {
		curStud := fmt.Sprintf("%v %v %v", val.Name, val.Age, val.Grade)
		result = strings.Join([]string{result, curStud}, "\n")
	}
	fmt.Println(result)
}
func main() {
	studStorage := storge.NewStorage()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		valid, stud := convertData(txt)
		if !valid {
			continue
		}
		studStorage.Put(stud)
	}
	printStudData(studStorage.Get())
}
