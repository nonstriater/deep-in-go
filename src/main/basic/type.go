package main

import "fmt"

func main()  {
	type MyInt1 int  //类型定义，定义一个新类型，与原类型(这里是int)是不同的两个类型，想要赋值的话，需要经过强制类型转换
	type MyInt2 = int  //类型别名(type aliases) 与原类型完全一样，可以相互赋值， golang 1.9 引入， 引入背景用于大型项目重构
	var i int =9
	//var i1 MyInt1 = i //编译不过
	var i2 MyInt2 = i
	//fmt.Println(i1)
	fmt.Println(i2)
}