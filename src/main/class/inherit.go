package main

import "fmt"

type People struct{
	identifer string
	name string
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People //类型嵌入， 跟继承还是有区别的
	major string //专业，比如是音乐
}

func (t *Teacher) ShowB() {
	t.name = "t name"
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA() // 会被强制类型转换？ 并不是真正意义的继承关系
}

//没有使用接口，通过方法重载实现了多态

/*
showA
showB
*/