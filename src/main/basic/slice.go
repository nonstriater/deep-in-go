package main

import "fmt"

func main()  {

	sli := make([]string, 10)
	sli = append(sli, "a")
	sli = append(sli, "b")
	sli = append(sli, "c")

	fmt.Printf("%#v", sli)
}