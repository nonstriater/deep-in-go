package main

func test(x int) (func(),func())  {
	return func() {
			println(x)
			x+=10
		}, func() {
			println(x)
		}
}

func main()  {
	a,b:=test(100)
	a()
	b()
}
/*
了解函数调用过程？
100
110
*/
