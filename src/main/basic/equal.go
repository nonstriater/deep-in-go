package main

import "fmt"

func main()  {

	if 10 == 10.0 {//true
		fmt.Println("true")
	}

	if 0.3*1 == 0.3 {//true
		fmt.Println("true")
	}

	if 0.1*3 == 0.3 {//true
		fmt.Println("0.1*3 == 0.3 true")
	}

	a := 0.1
	if a*3 == 0.3 {//false
		fmt.Println("a*3 == 0.3 true")
	}


}
