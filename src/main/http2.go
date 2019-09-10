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

	for {
		conn, err := lisener.Accept()//阻塞，获取来自客户端conn
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go echoFunc(conn) //goroutine 处理conn

	}
}

func echoFunc(c net.Conn)  {
	buf := make([]byte, 1024)

	for   {
		n, err := c.Read(buf) //读数据
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		c.Write(buf[:n])
	}
}
