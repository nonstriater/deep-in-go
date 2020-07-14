package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type ifoo interface{
	i_func1()
	i_func2()
}

type bar struct{}

func (b bar)i_func1()  {
	fmt.Println("hhh")
}

func (b bar)i_func2(a int)  {
	fmt.Println("lll")
}

type foo  struct{
	a int
	b string
}

func (o foo)fun1()  {
	fmt.Println(o.b)
}

func (o foo)Fun2(aa int)  {
	o.a = aa
}

func (o foo)Fun3(bb string)  {
	o.b = bb
}


func main() {
}

func Test_bacis(t *testing.T)  {
	str := "hello123567"
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.ValueOf(str))

	fmt.Println(reflect.TypeOf("kitty").Name())//string
	fmt.Println(reflect.TypeOf(123).Name())//int

	fmt.Println(reflect.TypeOf("kitty").String())//string
	fmt.Println(reflect.TypeOf(123).String())//int

	//fmt.Println(reflect.TypeOf("kitty").Bits())//crash
	fmt.Println(reflect.TypeOf(123).Bits())//64 bit(8个字节)
	fmt.Println(reflect.TypeOf(uint(123)).Bits())//64 bit(8个字节)
	fmt.Println(reflect.TypeOf(float32(123)).Bits())//32 bit(4个字节)

	//size
	fmt.Println(reflect.TypeOf("kitty").Size())//16字节
	fmt.Println(reflect.TypeOf("welcome a new world").Size())//16字节
	fmt.Println(reflect.TypeOf(123).Size())//8字节
	fmt.Println(reflect.TypeOf(float32(123)).Size())//4个字节
}

func Test_struct_basic(t *testing.T)  {
	f := foo{a:1, b:"hello"}
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.ValueOf(f))
}


func Test_struct(t *testing.T)  {
	f := foo{a:1, b:"hello"}
	typ := reflect.TypeOf(f)
	fmt.Println(typ.NumField())//2

	sf := typ.Field(0)
	fmt.Println(sf.Name)//a
	fmt.Println(sf.PkgPath)//command-line-arguments
	fmt.Println(sf.Type)//int
	fmt.Println(sf.Tag)// 空
	fmt.Println(sf.Offset)//0
	fmt.Println(sf.Index)//[0]
	fmt.Println(sf.Anonymous)//false

	sf = typ.Field(1)
	fmt.Println(sf.Name)//b
	fmt.Println(sf.PkgPath)//command-line-arguments
	fmt.Println(sf.Type)//string
	fmt.Println(sf.Tag)
	fmt.Println(sf.Offset)//8
	fmt.Println(sf.Index)//[1]
	fmt.Println(sf.Anonymous)//false

}

func Test_method(t *testing.T)  {

	f := foo{a:1, b:"hello"}
	typ := reflect.TypeOf(f)
	fmt.Println(typ.NumMethod())//2
	m := typ.Method(0)
	fmt.Println(m.Name)//Fun2
	fmt.Println(m.PkgPath)
	fmt.Println(m.Type) //func(main.foo, int)
	fmt.Println(m.Func) //0x10eacd0

	fmt.Println(m.Type.NumIn())//2, receiver是第一个参数
	fmt.Println(m.Type.NumOut())//0

	fmt.Println(m.Type.In(1))//int
}

func Test_kindAndElem(t *testing.T)  {

	f := foo{a:1, b:"hello"}
	typ := reflect.TypeOf(f)
	m := typ.Method(0)

	fmt.Println(typ.Kind())//struct
	fmt.Println(m.Type.Kind())//func

	//type 必须是 Array, Chan, Map, Ptr, or Slice 这几种类型之一
	mmap := map[string]int{"1":33,"2":55}
	fmt.Println(reflect.TypeOf(mmap))//map[string]int
	fmt.Println(reflect.TypeOf(mmap).Elem())//int 返回值是元素类型
}

func Test_implements(t *testing.T)  {
	f := foo{a:1, b:"hello"}
	typ := reflect.TypeOf(f)
	fmt.Println(typ.Implements(reflect.TypeOf(f)))//panic
}

func Test_align(t *testing.T)  {
	f := foo{a:1, b:"hello"}
	typ := reflect.TypeOf(f)
	fmt.Println(typ.Align())//8
	fmt.Println(typ.FieldAlign())//8
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////     Value  //////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


func Test_value_addr(t *testing.T)  {
	f := foo{a:1, b:"hello"}
	val := reflect.ValueOf(f)
	//fmt.Println(val.Addr())//crash
	fmt.Println(val.CanAddr())//false

	fo := &foo{a:1, b:"hello"}
	val = reflect.ValueOf(fo)
	//fmt.Println(val.Addr())//crash
	fmt.Println(val.CanAddr())//false
	fmt.Println(val.CanSet())//false

	ofo := val.Elem()//解引用
	fmt.Println(ofo.Addr())//&{1 hello}
	fmt.Println(ofo.CanAddr())//true
	fmt.Println(ofo.CanSet())//true

}

func Test_value_get_0(t *testing.T) {
	f := foo{a:1, b:"hello"}
	val := reflect.ValueOf(f)

	fmt.Println(val.Field(1))//hello
	fmt.Println(val.Field(0))//hello
}

func Test_value_get(t *testing.T) {
	f := &foo{a:1, b:"hello"}
	val := reflect.ValueOf(f)

	fmt.Println(val.Elem())//{1 hello}
	fmt.Println(val.Elem().Addr())//&{1 hello}
	//fmt.Println(val.Field(1))//
	//fmt.Println(val.Index(1))//第i个elem
}

func Test_value_set(t *testing.T) {

	f := &foo{a:1, b:"hello"}
	val := reflect.ValueOf(f)
	fm := val.Elem()

	a := fm.Field(0)
	if a.CanSet() {
		a.SetInt(22)
	}

	b := fm.Field(1)
	b.SetString("kitty")
	fmt.Println(fm)//{1 hello}

	//fmt.Println(val.Field(1))//
	//fmt.Println(val.Index(1))//第i个elem
}

func Test_value_call(t *testing.T){

}

func Test_value_interface(t *testing.T){

}

func Test_value_map(t *testing.T){

}

//go test -gcflags=-N -benchmem -test.count=3 -test.cpu=1 -test.benchtime=1s
func BenchmarkFunA(b *testing.B)  {
	num:=10
	for i:=0;i<b.N;i++{
		fmt.Sprintf("%d",num)
	}
}


