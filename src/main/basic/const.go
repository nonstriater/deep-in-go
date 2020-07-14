package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	o
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,o,p)
}

/*
0 1 zz zz zz 4
*/

