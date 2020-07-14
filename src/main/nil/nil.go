package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	AgeYears int
	Name string
	Friends []Person
}

func sayHi(p *Person) {
	fmt.Println("hi")
}

func (p *Person) sayHi() {
	fmt.Println("hi2")
}

func main(){

	interface1()
	string1()
	assign()
}

func test1()  {
	var p Person
	//if p != nil {//结构体类型变量不能 与 nil 进行比较
	//	fmt.Println("nil")
	//}

	fmt.Println(p)//{0  []}
	p.sayHi()//hi2

	var t *Person
	fmt.Println(t)//nil
	t.sayHi()//hi2
}

func interface1(){
	var i interface{}
	var p *int
	fmt.Println(i)
	fmt.Printf("%p\n",i)
	fmt.Printf("%p\n",p)
}

func string1()  {
	var s string
	//if s == nil{//不能对比
	//}

	fmt.Println(len(s))//0
	fmt.Println(unsafe.Sizeof(s))//16

}

func assign()  {
	nil := 123
	fmt.Println(nil) // 123

	// 如下代码行会产生编译错误, 因为当前作用域 nil 代表一个 nil 值.
	/*
	   var _ map[string]int = nil
	*/
}


