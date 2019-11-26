package main


type int1 interface {
	funcA()
}

type int2 interface {
	funcB()
}

type foo struct {

}

func (o foo)funcA()  {

}

func (o foo)funcB()  {

}


type bar struct {

}


func (b bar)funcA()  {

}

func main() {
	var d int1 = foo{}
}



