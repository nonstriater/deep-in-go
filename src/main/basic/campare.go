package main

import "fmt"

type data struct {
	num int
	fp float32
	complex complex64
	str string
	char rune
	yes bool
	events <-chan string
	handler interface{}
	ref *byte
	raw [10]byte
	//bytes []byte           //not comparable
	//doit func() bool       //not comparable
	//m map[string] string   //not comparable , 对比会编译报错
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",v1 == v2) //prints: v1 == v2: true
}

