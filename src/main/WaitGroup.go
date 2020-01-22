package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	//for j := 0; j < 10; j++ {
	//	go func() {
	//		fmt.Println("j: ", j)
	//		wg.Done()
	//	}()
	//}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("end")
}

/*
输出结果稳定：
i:  9
j:  10
j:  10
j:  10
j:  10
j:  10
j:  10
j:  10
j:  10
j:  10
j:  10
i:  0
i:  1
i:  2
i:  3
i:  4
i:  5
i:  6
i:  7
i:  8
end
*/