## Golang 设计模式

分为3类

- 创造型
- 结构型
- 行为型


## 创造型

- 单例
- 工厂

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



## 结构型

- facade
- delegate

## 行为型

- observer
- command(命令)
- 模板方法
- 状态机

