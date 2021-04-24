package _interface

type int1 interface {
	funcA()
}

type int2 interface {
	funcB()
}

type foo struct {

}

func (o foo)talk(v interface{})  {
	if v == nil {
		println("nil")
	} else {
		println("non-nil")
	}
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
	var b bar
	NilOrNot(b)


}

func NilOrNot(v interface{}) {
	if v == nil {
		println("nil")
	} else {
		println("non-nil")
	}
}




