package main

import "fmt"

func main()  {

	testMap()
}

func testMap()  {
	json := map[string]int{"1":11, "2":22}
	fmt.Println(json)
}

func testMap2()  {
	json := map[int]int{1:11, 2:22}
	fmt.Println(json)
}

func testMap3()  {
	json := map[byte]int{'a':11, 'b':22}
	fmt.Println(json)
}

