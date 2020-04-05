package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/a1", a1)//设置路由
	http.ListenAndServe(":8001",nil)

	fmt.Println("exit..")

}

func a1(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("hello"))
}