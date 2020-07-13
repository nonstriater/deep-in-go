# deep-in-go

- array 数组
- slice/map "引用类型"
- for/range 循环遍历
- const 常量
- defer 延迟
- panic/recover 异常处理
- type  类型定义&类型别名
- func   函数调用过程
- struct  结构体
- nil 零值处理
- reflect 反射
- inherit  "继承" 类型嵌入, 与继承有哪些区别？
- interface  接口，泛型，鸭子类型，动态调用
- chan 管道
- select/case 
- 锁: Mutex/RWMutex/Once
- WaitGroup 等待组 协程同步
- goroutine 协程调度器
- memory 内存模型
- gc 垃圾回收
- pprof  性能
- timer  定时器
- http  网络模型 client/server
- Go 进程初始化过程

## 简介

学习一门语言，需要对每个特性：

- 写 Demo 进行各种特性验证，有什么坑之类的 
- 为什么需要这个语法特性，语言设计者是如何考虑的？与其他语言如 Golang/Python/Java/OC/JS 等横向对比
- 查看编译过程，了解语法特性的编译时过程
- 查看运行过程，看看 runtime 对语法特征的支持
- 看源码，golang是开源的
- 了解语法特性的各种使用场景；这个需要有工程经验，也可以阅读优秀的开源代码获得经验
- 设计模式，了解常用的设计模式在这门语言下如何实现，如单例,工厂,代理,facade,观察者,模板方法等
- 记笔记，总结

查看编译过程方法 

```
go tool 6g -S hello.go  // -S  从源码直接生成汇编代码
go tool compile -S -N -l main.go  //汇编代码
```

Goland IDE 上怎么查看汇编代码？

暂未找到

查看运行时过程的方法 

```
runtime
```



## array 数组

```
arr1 := [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3} //编译期间进行大小推断
[]interface{}
```

不可变数组分配在栈区，函数参数传递时采用的值拷贝。

## slice/map "引用类型"

slice/map 属于"引用类型"，容器内的数据在运行时存储在堆上

```
slice := []int{1, 2, 3} //字面量初始化
slice := make([]int, 10)
```

```
hash := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}
```

## for/range 循环遍历

`for .. range <容器>` 用来枚举 `数组/slice/map/channel` 容器结构; 如下：

```
func main() {
	arr := []int{1, 2, 3}
	for i, _ := range arr {
		println(i)
	}
}
```

这里主要介绍下容易踩的坑

```
func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
		{Name: "ran", Age: 25},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu //for/range 里取地址 都是同一个地址值？为什么？总是最后一个元素的地址
		fmt.Printf("%p,%v",&stu,stu)
		fmt.Println()
	}
	fmt.Println(m)
	fmt.Println(m["zhou"])
	fmt.Println(m["li"])
	fmt.Println(m["wang"])
}

/*
0xc0000b0000
0xc0000b0000
0xc0000b0000
map[li:0xc0000b0000 wang:0xc0000b0000 zhou:0xc0000b0000]
&{wang 22}
&{wang 22}
&{wang 22}
*/
```


## const 常量

`iota` 关键字在 const 关键字出现时初始化为 0, const 中每新增一行常量声明将使 iota 计数一次.
因此如下常量 `p = 4`

```
const (
   x = iota 
   y
   z = "zz"
   k
   p = iota
)

func main()  {
   fmt.Println(x,y,z,k,p)
}

/*
0 1 zz zz 4
*/
```

## defer延迟

defer关键词是一个语法糖，需要编译器和运行时共同之处才能实现。

主要用于操作必须要成对出现的场景，比如对象创建和释放，加锁和释放锁。
释放资源的操作往往在用完之后容易忘记，造成资源泄露或资源不可再次获得；

因此使用 defer 语句可以
将资源的创建和释放代码写在一起，但是资源释放发生在函数

函数中出现多个defer时，会逆序执行，因为在运行时是一个栈式结构； defer也可以嵌套，执行循序由外向内，
这些都比较容易理解，但是当defer和局部变量和返回值等混在一起就不太容易搞明白

```
func main() {
    func_b_0()
    func_b_1()
    func_b_2()
}
func func_b_0() {
    fmt.Println("func_b_0...top")
    a := 5
    defer fmt.Println("a=", a)//打印值: 5, 原因: 在a := 5之后, 这句话就已经编译好, 只是放到最后才执行, 所以为5
    a++
}

//defer 函数之后还有可能修改变量，造成defer函数中不确定性，应该避免这种写法
func func_b_1() {
    fmt.Println("func_b_1...top")
    a := 5
    defer func() {
        fmt.Println("a=", a) //打印值: 6, 原因: 在a := 5之后 , 此句就已经编译好, 但没有执行, 放到最后才执行, 最后执行时才传入当时的参数6
    }()
    a++
}

//如果defer函数中使用局部变量，最好使用这种方式
func func_b_2() {
    fmt.Println("func_b_2...top")
    a := 5
    defer func(a int) {
        fmt.Println("a=", a)//打印值: 5, 原因: 在a := 5之后, 此句话就已经编译好, 并传入了当时的参数5, 只是放到了最后才执行
    }(a)
    a++
}
```


## panic/recover 异常处理

recover() 必须在 defer 中调用， 当 panic 发生时，会编译本协程的defer链，如果defer中存在recover()语句就可以捕获panic

```
func defer_call2()  {
   defer func() {
      if ok := recover(); ok != nil {
         fmt.Println("recover," + ok.(string))//recover 返回值就是 panic()传入的参数
      }
   }()
   panic("error")
}

//recover,error
```

## type 类型定义&类型别名

- 类型定义： 定义一个新类型，与原类型(这里是int)是不同的两个类型，想要赋值的话，需要经过强制类型转换
- 类型别名(type aliases): 与原类型完全一样，可以相互赋值， golang 1.9 引入， 引入背景用于大型项目重构

```
func main()  {
   type MyInt1 int  //类型定义，定义一个新类型
   type MyInt2 = int  //类型别名(type aliases) 与原类型完全一样，可以相互赋值
   var i int =9
   //var i1 MyInt1 = i //编译不过
   var i2 MyInt2 = i
   //fmt.Println(i1)
   fmt.Println(i2)
}
```


## func 函数调用过程

- 函数调用过程参数，返回值通过栈传递，因此函数可以由多个返回值；C语言中返回值用 EAX/R0 寄存器存储
- 归属于struct的方法调用，函数调用过程是把struct作为对一个参数来完成函数调用过程
- 一个函数栈完成,参数/局部变量会释放而触发gc

重点说下 slice/map 作为参数/返回值时函数调用过程, slice/map 属于"引用类型"，函数调用作为参数/返回值时，
传递的仅仅是指针地址，容器内的数据在运行时存储在堆上。

```
//slice类型返回值，返回的仅仅是指针地址
func getApis() []string {
	var apis []string
	apis = append(apis, "1")
	apis = append(apis, "2")
	
	return apis
}

//slice类型作为参数，传递的仅仅是指针拷贝
func startIPC(apis []string) {
	fmt.Printf("%v", apis)
}
```


## struct 结构体

Go 中使用 struct 实现了面向对象编程，用 struct 封装一个类类型， 然后可以给类型定义自有的方法

面向对象的语言主要在大型系统的设计上，有较好的代码组织结构关系(拆分好的模块以及之间的关系)，以及代码良好的复用性和灵活性；
我们看一个语言是不是面向对象语言，看3个方面：

- 封装性
- 继承，主要为代码复用， Go 通过类型嵌入的方式实现
- 多态，增强灵活性，通过 interface 接口实现

## inherit  "继承" 类型嵌入, 与继承有哪些区别？

```
type People struct{}
type Teacher struct {
   People
}
```

类型嵌入可以继承基类的变量和方法。但并不支持重载方法实现多态，因此实现多态就只有另一条路：interface接口


## interface 接口，泛型，鸭子类型，动态调用

接口定义如下：

```
type I interface {
    Get() int
}
```

接口作为函数参数/返回值 时的函数调用过程？空的interface{} 有什么坑？interface 实现原理是什么？

```
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

//output: non=nil
```




多态和动态调用还不太一样。



## nil 零值处理

介绍几个 nil 零值处理的坑

- 指针，slice/map/channel,function,interface 零值都是nil;  struct 类型的零值并不是 nil，对比的话编译报错
- 值为nil的指针解引用会崩溃
- `nil==nil` 不能比较，会编译报错
- nil 标识符可以被覆盖，即可以这些写 `nil := 123`，但不要这样做


## reflect 反射

很多框架都用到reflect 实现一些动态的功能，比如 gorm。

反射具体使用参考 `reflect.Type` 和 `reflect.Value` 的 api

简单介绍下反射的基本实现原理。



## chan 管道

chan 信道是 goroutine 通讯的机制；

类似一个消息队列，在这个队列写满的情况继续写就会阻塞当前goroutine，同理，空的情况下读 chan 信道 也会阻塞当前 goroutine, 这个也常被用来实现多协程场景下的协程同步。

chan 初始化和使用

```
var channel chan int = make(chan int)
//或
channel := make(chan int)

data :=  <- channel  //从 chan 读取数据
channel <-   data   //数据写入 chan
```

chan 信道读取和发生都会造成阻塞  
必须一个goroutine读，一个goroutine写； 在单一的信道里 读或者写 chan都会导致死锁


## select/case 

用来实现多个信道的等待操作，加一个default操作，可以实现非阻塞多路 `select`

## 锁: Mutex/RWMutex/Once


## WaitGroup 等待组 


## goroutine 协程


## memory 内存模型


## gc 垃圾回收


## pprof  性能


## timer  定时器


## http  网络模型 client/server


## Go 进程初始化过程


# 参考

https://draveness.me/golang/

