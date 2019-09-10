package main

import "net/http"

func main() {
	http.HandleFunc("/a1", a1)//设置路由
	http.ListenAndServe(":8001",nil)
}

func a1(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("hello"))
}