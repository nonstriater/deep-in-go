package main
import "fmt"

func main() {
	arr1()
	arr2()
	arr3()
	arr4()
	arr5()
}

func arr1()  {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func arr2(){
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func arr3()  {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 5)
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func arr4()  {
	str := "world"
	end := append([]byte("hello "), str...)
	fmt.Println(end)
}

//[]byte 与 string 相互转换
func arr5()  {
	str := "helllo"//不可变
	bytes := []byte(str) //string to []byte
	println(bytes)

	str2 := []byte{'a','b','c'}//可变
	str3 := string(str2) //[]byte to string
	println(str3)
}