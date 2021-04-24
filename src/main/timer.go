package main

import "time"

//CPU 负载高？没发现有
func main()  {

	for i := 0;  i < 800 ; i++   {
		go test1()
	}

	//block here
	test1()
}

func test1()  {
	var ticker = time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()
	var counter = 0
	for {
		select {
		case <-ticker.C:
			counter += 1
		}
	}
}