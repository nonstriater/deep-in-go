package main
import "fmt"

func main() {
	arr1()
	arr2()
	arr3()
	arr4()
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

