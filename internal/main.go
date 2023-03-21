package main

import (
	"bufio"
	"fmt"
	"learning/module27/pkg/storge"
	"os"
	"strconv"
	"strings"
)

type Storage interface {
	Put(string, int, int)
	Get() string
}

func convertData(data string) (bool, string, int, int) {
	dataSplited := strings.Split(data, " ")
	if len(dataSplited) != 3 {
		fmt.Println("Incorrect input")
		return false, "", 0, 0
	}
	name := dataSplited[0]
	age, err := strconv.Atoi(dataSplited[1])
	if err != nil {
		fmt.Println("Incorrect input", err)
		return false, "", 0, 0
	}
	grade, err := strconv.Atoi(dataSplited[2])
	if err != nil {
		fmt.Println("Incorrect input", err)
		return false, "", 0, 0
	}
	return true, name, age, grade
}

func main() {
	studStorage := storge.NewStorage()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		valid, name, age, grade := convertData(txt)
		if !valid {
			continue
		}
		studStorage.Put(name, age, grade)
	}
	fmt.Println(studStorage.Get())
}
