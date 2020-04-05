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
- channel  管道
- select/case 
- lock 锁, Mutex/RWMutex/Once
- waitgroup  协程同步：等待组
- goroutine 协程
- memory 内存模型
- gc 垃圾回收
- pprof  性能
- timer  定时器
- http  网络模型 client/server

对于每个部分，简单介绍下使用以及需要注意的一些坑。

## array 数组

```
arr1 := [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3} //编译期间进行大小推断
[]interface{}
```

不可变数组分配在栈区，函数参数传递时采用的值拷贝。

## slice/map "引用类型"

slice/map 属于"引用类型"，容器内的数据在运行时存储在堆上

```` 
slice := []int{1, 2, 3} //字面量初始化
slice := make([]int, 10)
````

```
hash := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}
```

## for/range 循环遍历

用来枚举 slice/map/channel


## const 常量

iota 在 const 关键字出现时初始化为 0, const 中每新增一行常量声明将使 iota 计数一次.
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



## panic/recover 异常处理

## type  类型定义&类型别名

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

## nil 零值处理



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

## reflect 反射

## struct 结构体

## inherit  "继承" 类型嵌入, 与继承有哪些区别？
## interface  接口，泛型，鸭子类型，动态调用
## channel  管道
## select/case 
## lock 锁, Mutex/RWMutex/Once
## waitgroup  协程同步：等待组
## goroutine 协程
## memory 内存模型
## gc 垃圾回收
## pprof  性能
## timer  定时器
## http  网络模型 client/server




# 参考

https://draveness.me/golang/

