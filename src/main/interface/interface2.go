package _interface

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo(x)

	//var s string = nil //nil不能赋值个string
	var s string
	Foo(s)
}

