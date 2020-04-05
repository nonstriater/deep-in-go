package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func test(x int) (func(),func())  {
	return func() {
			println(x)
			x+=10
		}, func() {
			println(x)
		}
}

func main()  {
	a,b:=test(100)
	a()
	b()
}
/*
了解函数调用过程？
100
110
*/

func getApis() []string {
	var apis []string
	apis = append(apis, "1")
	apis = append(apis, "2")

	return apis
}

func startIPC(apis []string) {
	fmt.Printf("%v", apis)
}

