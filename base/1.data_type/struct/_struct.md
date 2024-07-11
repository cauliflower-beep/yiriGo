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

具体可参看本小结[代码](.\5_compare.go)。