package main

import (
	"fmt"
)

func executeFn(fn func() int) int {
	return fn();
}

func main() {
	a := 1
	b := 2
	c := executeFn(func() int {
		a += b
		return a
	})
	fmt.Printf("%d %d %d\n", a, b, c) //3 2 3
}

//调用闭包时参数并不通过栈传递, 而是通过寄存器rdx传递
//通过查看汇编代码，闭包传递的内容中，a 变量传递使用的是 LEAQ 指令，传递的 a 的地址