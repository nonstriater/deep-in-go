package main

import (
"net"
"fmt"
)

func main()  {
	lisener,err :=  net.Listen("tcp", ":8888")//server监听
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("listen at 8888...")

	for {
		conn, err := lisener.Accept()//阻塞，获取来自客户端conn
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go echoFunc(conn) //goroutine 处理conn

	}
}

//为什么有的时候会执行两遍？
func echoFunc(c net.Conn)  {
	fmt.Println("receive conn")

	//调用同步阻塞IO api



}
