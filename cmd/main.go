package main

import (
	"bufio"
	"learning/module27/pkg/storge"
	"os"
)

func main() {
	studStorage := storge.NewStorage()

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		storge.Put(txt, studStorage)
	}
	storge.Get(studStorage)
}
