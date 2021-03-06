# Go学习

* [Go基础](#GO-BASIC)

### <a name="GO-BASIC"></a>Go基础知识

+ Go的第一个参数默认的是当前的执行路径

+ 定义参数
````
var power int = 9000;
power := 9000;
// := 被用来声明变量并且赋值（一个变量不能被声明两次）。
````

+ = 和:= 的区别
> = 号只有赋值操作，:= 先声明变量再赋值
> + var NAME TYPE 声明一个变量并且赋值为零
> + NAME := VALUE 声明并且赋值
> + NAME = VALUE 赋值

+ 函数定义
```$xslt
func log(message string){
}
func add(a int, b int) int{
}
func power(name string) (int, bool){
}
```

+ 定义结构
```$xslt
type Saiyan struct{
    Name string
    Power int
}
```
+ &以及*

>+ & 取地址符，后面跟的是变量
>+ \* 取值,后面跟的必须时一个地址,表明变量是一个指针

+ 构造
> construct
```$xslt
func NewSaiyan (name string, power int) * Saiyan{
    return  &Saiyan{
        Name: string,
        Power: power,
    }
 }
```
+ New
```$xslt
goku := new(string)
// same as
goku := &Saiyan{ }
```

+ Composition(组合)
> 组合优于继承

+ 值还是指针

|~|值|指针|
|:-----:|:-----:|:-----:|
|场景|局部变量，域，返回值，函数参数值，函数接受值| 不需要副本的时候|

+ Arrays
```$xslt
scores := []int{2,23,4,5}
```

+ Slices
```$xslt
scores := make([]int, 10)
```
+ Go是以二倍的当时增长

+ Map

+ Error Handling 错误处理









