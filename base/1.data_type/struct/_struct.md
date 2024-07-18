## 概述

go中的 `struct` 属于复合类型，并非引用类型；与其他面向对象编程语言中的 class 类似，struct可以定义字段（属性）和方法。

点击查看[小节代码](.\struct.go)。

### 定义

结构体有如下定义规则：

1. 同类型的字段可以放在一行定义，但比较好的编程习惯是分行定义；
2. struct可以通过复合其他结构体来构建更复杂的结构体，但不能包含自身；
3. 只提供类型，而不写字段名的方式，称之为**匿名字段**，也叫嵌入字段；匿名字段可以是其他struct、自定义类型；

### 值传递

结构体与数组一样，都是值传递；意味着当把数组或结构体作为实参传给函数的形参时，会复制一个副本。所以为了提高性能，一般不会把数组直接传递给函数，而是使用切片(引用类型)代替，而把结构体传给函数时，可以使用指针结构体。

数据量小的结构体，传拷贝比传指针快，但如果结构体里面有map，传指针一定比传拷贝要快。

### tag

定义结构体字段时，除字段名称和数据类型外，还可以使用`反引号`为结构体字段声明元信息。这种元信息称为Tag，用于编译阶段关联到字段当中。

Tag由反引号括起来的一系列用空格分隔的 key:"value" 组成。

### 方法

Go语言中，将函数绑定到具体的类型中，则称该函数是该类型的方法。其定义的方式是在func与函数名称之间加上具体类型变量，这个类型变量称为方法接收器。

并不是只有结构体才能绑定方法，任何类型都可以绑定方法。

方法分为值接收器和指针接收器，因为`struct`是值传递，所以我们指定结构体为方法接收器时，通常传入结构体指针，否则函数操作的只是结构体的一个副本，并不会对原结构体造成影响。

## 空结构体

不包含任何字段的结构体称为空结构体，struct{}表示一个空的结构体。直接定义一个空结构体并没有意义。

空结构体有些常见用法：

1. 作为仅有方法的结构体；

2. 构建 set；

   go中没有set，如何利用map来实现一个set？map和set是两个抽象的数据结构，map存储一个键值对集合，其中键不重复，set存储一个不重复的元素集合。

   本质上set可以视为一种特殊的map，set其实就是map中的键。

   用map模拟一个set，就要把值置为struct{}。struct{}本身是不占任何空间的，可以避免任何多余的内存分配；

3. 并发编程中，channel之间的通讯可以使用一个struct{}作为信号量，channel <- struct{}{}，也是为了节省空间。

## 结构体的比较

结构体有如下比较规则：

1. 结构体能否**直接比较**，与字段类型、字段值、字段个数、字段顺序相关，

   如下代码：

   ```go
   sn1 := struct{
       age int
       name string
   }{age:16,name:"conan"}
   sn2 := struct{
       name string
       age int
   }{age:16,name:"conan"}
   ```

   sn1、sn2字段顺序不同，不能直接比较，但可借助`reflect.DeepEqual`进行比较：

   ```go
   if reflect.DeepEqual(sn1, sn2) { // false 因为字段顺序不同
   		fmt.Println("sn1 == sn2")
   	} else {
   		fmt.Println("sn1 != sn2") // 比较结果
   	}
   ```

   再如：

   ```go
   sn3 := struct{
       age int
       name string
   }{age:16,name:"conan"}
   ```

   则可以直接比较sn1与sn3：

   ```go
   fmt.Println(sn1 == sn3) // true
   ```

2. `map`、`slice`作为结构体的字段类型时，不能直接比较，但同样可以使用`reflect.DeepEqual`进行比较：

   ```go
   var sm1 = struct{
       name string
       grade map[string]int
   }{name:"tom",grade:map[string]int{"math":92,"english":93}}
   
   var sm2 = struct{
       name string
       grade map[string]int
   }{name:"tom",grade:map[string]int{"math":92,"english":93}}
   ...
   if reflect.DeepEqual(sm1,sm2){
       fmt.Println("sm1 == sm2")
   }else{
       fmt.Println("sm1 != sm2")
   }
   ```

具体可参看本小结[代码](.\compare.go)。