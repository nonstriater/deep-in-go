package main

func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
	println(DeferFunc4(1))
}

func DeferFunc1(i int) (t int) {
	t = i //这行代码为啥能编译通过？没有 var 关键字申明， t 是返回值
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4(i int) (t int) {
	t += i
	return 2
}

/*
4
1
3
2
*/
