package main

import (
	"fmt"
	"net"
	"strings"
)

func main()  {
	lisener,err :=  net.Listen("tcp", ":8888")//server监听
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("listen at 8888...")

	for {
		//阻塞，获取来自客户端conn
		//客户端telnet localhost 8888 连接
		conn, err := lisener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		//goroutine 处理conn
		fmt.Println("Connection from " , conn.RemoteAddr())
		go blockFunc(conn)  // golang 网络架构是 goroutine + 同步阻塞IO 架构模式

		//golang中多路复用用的是netpoll，是基于epoll/kqueue/iocp的封装
		//go multiplexingFunc(conn) //goroutine + 多路复用

		//go nonblockFunc(conn) //同步非阻塞IO
		//go asyncIOFunc(conn)//异步非阻塞
	}
}

//为什么有的时候会执行两遍？
func blockFunc(c net.Conn)  {
	if c == nil {
		return
	}
	defer c.Close()

	buf := make([]byte, 4096)

	for  {
		//底层实现，仍然是用的 netpoll 多路复用
		count, err := c.Read(buf) //阻塞IO， 可以设置read timeout
		if err != nil && count == 0{
			println("read err or empty:" + err.Error())
			c.Close()
			return
		}

		readStr := (string(buf[0:count]))//eg: ping\r\n
		readStr = strings.TrimSpace(readStr)//去掉空格符号
		fmt.Println("read message:  " +  readStr)

		switch readStr {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			c.Write([]byte("echo\n"))
		case "quit":
			c.Close()
		default:
			fmt.Println("unknown cmd:  " + readStr)
		}
	}


}

func nonblockFunc(c net.Conn)  {
	fmt.Println("receive conn")


}

func multiplexingFunc(c net.Conn){

}

func asyncIOFunc(c net.Conn){

}