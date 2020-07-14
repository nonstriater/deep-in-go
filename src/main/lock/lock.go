package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"sync"
)

/*
chan
Mutex
RWMutex
Once
WaitGroup
*/

func main(){
	mutex := sync.Mutex{}
	fmt.Printf("mutex: %v", mutex)
}

//当结构类型包含 sync.Mutex 或者同步这种字段时，必须使用指针以避免成员拷贝
//mutex 对象 发送拷贝会怎样？


