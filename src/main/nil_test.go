package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_nil_campare(t *testing.T)  {
	//fmt.Println(nil == nil)

	//var m map[int]string
	//var ptr *int
	//fmt.Printf(m == ptr)
}

func Test_nil_address(t *testing.T)  {
	var m map[int]string
	var ptr *int
	var l []int
	fmt.Printf("%p\n", m)//0x0
	fmt.Printf("%p\n", ptr)//0x0
	fmt.Printf("%p", l)//0x0
}

func Test_nil_len(t *testing.T){
	var s string
	//if s == nil{//不能对比
	//}
	fmt.Println(len(s))//0
	fmt.Println(unsafe.Sizeof(s))//16, 2个机器字长

	var l []int
	fmt.Println(unsafe.Sizeof(l))//24, 4个机器字长？
}
