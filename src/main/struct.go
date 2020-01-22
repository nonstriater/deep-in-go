package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main()  {
	pase_student()
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu //这里取地址都是同一个地址值？为什么？
		fmt.Printf("%p",&stu)
		fmt.Println()
	}
	fmt.Println(m)
	fmt.Println(m["zhou"])
	fmt.Println(m["li"])
	fmt.Println(m["wang"])
}

/*
0xc0000b0000
0xc0000b0000
0xc0000b0000
map[li:0xc0000b0000 wang:0xc0000b0000 zhou:0xc0000b0000]
&{wang 22}
&{wang 22}
&{wang 22}
*/
