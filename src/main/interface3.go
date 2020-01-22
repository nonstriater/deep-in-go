package main

import (
	"fmt"
)

type People2 interface {
	Show()
}

type Student struct{}
func (stu *Student) Show() {

}

func live() People2 {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/*
BBBBBBB
*/

