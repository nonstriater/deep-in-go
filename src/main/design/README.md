## Golang 设计模式

分为3类

- 创造型
- 结构型
- 行为型


## 创造型

- 单例
- 简单工厂  factory/builder
- 工程方法 factory method
- 抽象工程 abstract factory
- 原型模式 prototype


### 单例

golang 中使用全局变量，最常见的GORM DB 如下：

```
var db *gorm.DB
func SetGormDB(gdb *gorm.DB) {
	db = gdb
}

func GetGormDB() *gorm.DB {
	return db
}
```

单例中 Lazy mode (不带锁)

```
type singleton struct {
}

// private
var instance *singleton

// public
func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{}     // not thread safe
    }
    return instance
}
```

带锁的单例

```
import (
    "sync"
)
 
type singleton struct {
}
 
var instance *singleton
var once sync.Once
 
func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```


全局变量 和 单例 有什么优劣势？

全局变量一般在进程初始化时候赋值，没有lazy load；且运行过程中变量值可能改变
在多线程情况下访问需要手动加锁


### 工厂模式 Factory

golang 中 为什么抛弃了  constructor? 
java 里的 Factory Pattern, Builder Patter 还需要吗？



## 结构型

- facade
- delegate(proxy)
- adapter 适配器 
- bridge 桥接


## 行为型

- observer
- template 模板方法
- strategy 策略
- command(命令)
- chain 责任链
- state 状态机


