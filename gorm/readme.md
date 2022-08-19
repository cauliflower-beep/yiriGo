## 优缺点

**在项目中使用orm的好处很多：**

1. 防止直接拼接sql语句引入sql注入漏洞
2. 方便对modle进行统一管理
3. 专注业务，加速开发

**坏处也是显而易见的:**

1. 开发者与最终的sql语句隔了一层orm，因此可能会不慎引入烂sql
2. 依赖于orm的成熟度，无法进行一些「复杂」的查询。当然，复杂的查询一大半都是应该从设计上规避的

## 几个注意点

1. **留意不合法的时间值**

   MySQL的`DATE/DATATIME`类型可以对应Golang的`time.Time`。但是，如果`DATE/DATATIME`不慎插入了一个无效值，例如2016-00-00 00:00:00, 那么这条记录是无法查询出来的。会返回`gorm.RecordNotFound`类型错误。零值0000-00-00 00:00:00是有效值，不影响正常查询。

2. **留意tagsql:"default:null"**

   gorm对各种tag的支持非常完善。但是有些行为跟直觉不太一致，需要注意。当对某字段设置`tagsql:"default:null"`时，你想通过`update`设置该字段为`null`就不可能了，只能通过`raw sql`。这是gorm设计的取向问题。

3. **如何通过gorm设置字段为null值**

   字段允许为`null`值肯定是设计存在问题。但是，往往前人埋下的坑需要你去填。gorm作者给出了两种方法，以`string`为例：

   在golang中，声明该字段为`*string`

   使用`sql.NullString`类型

   推荐使用后者。

4. **留意连接串中的loc**

   例如通过如下连接串打开mysql连接：

   ```go
   db, err := gorm.Open("mysql", "db:dbadmin@tcp(127.0.0.1:3306)/foo?charset=utf8&parseTime=true&loc=Local")
   ```

   `parseTime=true&loc=Local`说明会解析时间，时区是机器的local时区。

   机器之间的时区可能不一致会设置有问题，这导致从相同库的不同实例查询出来的结果可能解析以后就不一样。因此推荐将loc统一设置为一个时区，如`parseTime=true&loc=America%2FChicago

## 结构体标签

​	使用gorm进行数据的增删改查操作，每个表或每个数据结果都会对应一个结构体。

​	有时在使用这个结构体时，会额外增加一部分字段，为一些其他业务逻辑所使用。这种情况如何有效避免一些字段不被gorm进行处理呢，整理部分常用结构体字段供查阅：

| 标签名             | 作用                                           | 案例                                    |
| ------------------ | ---------------------------------------------- | --------------------------------------- |
| column             | 指定字段名，如果不指定通常是字段的小写驼峰格式 | gorm:"column:usr_id;"                   |
| primaryKey         | 是否是主键                                     | gorm:"column:usr_id;primaryKey"         |
| unique             | 是否唯一                                       | gorm:"column:usr_id;unique"             |
| default            | 指定默认值                                     | gorm:"column:usr_id;default:123"        |
| -                  | 忽略该字段                                     | gorm:"-"                                |
| ->                 | 只读                                           | gorm:"->;column:usr_id"                 |
| <-                 | 允许读和写                                     | gorm:"<-;column:usr_id"                 |
| <-:update          | 允许读和更新                                   | gorm:"<-:update;column:usr_id"          |
| <-:create          | 允许读和创建                                   | gorm:"<-:create;column:usr_id"          |
| ->:false;<-:create | 仅创建                                         | gorm:"->:false;<-:create;column:usr_id" |

## 技巧

例子表：

| name  | number | age  |
| ----- | ------ | ---- |
| goku  | 1      | 33   |
| lufy  | 2      | 23   |
| conan | 3      | 13   |

- 查看gorm生成的sql语句

  ```go
  stmt := db.Session(&gorm.Session{DryRun:true}).Where("id = ?",1).First(&User{}).Statement 
  sql := stmt.SQL.String()
  println(sql)
  ```

  

- 