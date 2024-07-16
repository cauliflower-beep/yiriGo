## 概念

项目开发中，可能会随时调整声明的模型，比如添加字段和索引。

使用gorm的自动迁移功能，可以始终让数据库表结构保持最新状态。

此外，gorm提供了一些迁移接口的方法，方便操作数据库表、字段和索引。

## AutoMigrate

AutoMigrate 用于自动迁移 schema（模式），保持 schema 为最新。

**`注意`** AutoMigrate 会创建表、缺失的外键、约束、列和索引。 并且会更改现有列的类型，如果大小、精度、是否为空可以更改。 但不会删除未使用的列，以保护您的数据。(只增不减)

在执行 AutoMigrate时，我们需要先声明模型。

```go
type User struct {
	gorm.Model
	Name string
	Age  uint
}

type Product struct {
	gorm.Model
	Name  string
	Price int
}

type Order struct {
	gorm.Model
	UserID    int
	ProductID int
}
```

建立数据库连接后，执行 AutoMigrate：

```go
var db *gorm.DB

func init() {
	var err error
	//我这里用到数据库是mysql，需要配置DSN属性[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8&parseTime=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&User{}, &Product{}, &Order{})
}
```

## Migrator

GORM 提供了 Migrator 接口，该接口为每个数据库提供了统一的 API 接口，可用来为您的数据库构建独立迁移，例如：

SQLite 不支持 ALTER COLUMN、DROP COLUMN，当你试图修改表结构，GORM 将创建一个新表、复制所有数据、删除旧表、重命名新表。

一些版本的 MySQL 不支持 rename 列，索引。`GORM 将基于使用 MySQL 的版本执行不同 SQL。`

```go
type Migrator interface {
  // AutoMigrate
  AutoMigrate(dst ...interface{}) error

  // Database
  CurrentDatabase() string
  FullDataTypeOf(*schema.Field) clause.Expr

  // Tables
  CreateTable(dst ...interface{}) error
  DropTable(dst ...interface{}) error
  HasTable(dst interface{}) bool
  RenameTable(oldName, newName interface{}) error
  GetTables() (tableList []string, err error)

  // Columns
  AddColumn(dst interface{}, field string) error
  DropColumn(dst interface{}, field string) error
  AlterColumn(dst interface{}, field string) error
  MigrateColumn(dst interface{}, field *schema.Field, columnType ColumnType) error
  HasColumn(dst interface{}, field string) bool
  RenameColumn(dst interface{}, oldName, field string) error
  ColumnTypes(dst interface{}) ([]ColumnType, error)

  // Constraints
  CreateConstraint(dst interface{}, name string) error
  DropConstraint(dst interface{}, name string) error
  HasConstraint(dst interface{}, name string) bool

  // Indexes
  CreateIndex(dst interface{}, name string) error
  DropIndex(dst interface{}, name string) error
  HasIndex(dst interface{}, name string) bool
  RenameIndex(dst interface{}, oldName, newName string) error
}
```

### 数据库接口

CurrentDatabase 返回当前使用的`数据库名`

```go
db.Migrator().CurrentDatabase()
```

### 数据表接口

操作数据库表，必须先声明模型。

1. CreateTable 创建数据表

   ```go
   err := db.Migrator().CreateTable(&User{})
   if err != nil {
     fmt.Printf("创建数据库表失败，错误:%s\n", err)
     return
   }
   fmt.Println("创建数据库表成功")
   ```

   默认情况下，GORM 会约定使用 ID 作为表的主键，可以通过标签 `gorm:"primarykey"` 将其它字段设为主键。
   通过将多个字段设为主键，以达到创建复合主键，整型字段设为主键，默认为启用 AutoIncrement，如果需要禁用，使用标签 `autoIncrement:false`。

   GORM 约定使用结构体名的`复数形式作为表名`，不过也可以根据需求修改，可以实现Tabler 接口来更改默认表名，不过这种方式不支持动态变化，它会被缓存下来以便后续使用，如果想要使用动态表名，可以使用Scopes.
   GORM 约定使用结构体的字段名作为数据表的字段名,默认GORM 对 struct 字段名使用`Snake Case`命名风格转换成 MySQL 表字段名(需要转换成小写字母)，也可以通过标签 column 修改。

2. HasTable 检查对应的数据表是否存在

   ```go
   isExist := db.Migrator().HasTable(&User{})
   //isExist := db.Migrator().HasTable("users")
   if !isExist {
   	fmt.Printf("users 表不存在\n")
   	return
   }
   fmt.Printf("users 表存在\n")
   ```

3. DropTable 如果存在表则删除（删除时会忽略、删除外键约束)

   ```go
   db.Migrator().DropTable(&User{})
   // err := db.Migrator().DropTable("users")
   fmt.Printf("users 表删除成功\n")
   ```

4. RenameTable 重命名表

   ```go
   // db.Migrator().RenameTable("users", "user_infos")
   //若是users表存在则改名为user_infos表，反之亦然
   if b := db.Migrator().HasTable(&User{}); b {
   	db.Migrator().RenameTable(&User{}, &UserInfo{})
   	fmt.Printf("users 表名修改成功\n")
   } else {
   	db.Migrator().RenameTable(&UserInfo{}, &User{})
   	fmt.Printf("user_infos 表名修改成功\n")
   }
   // 个人推荐用结构体模型来进行以上操作，数据库的结构可以统一固定，这也是迁移的目的。
   ```

### 数据表字段接口

操作数据库表字段，必须先声明模型。

1. AddColumn 添加字段

   `注意`

   - 必须先声明模型。
   - 数据表不存在的字段名，且结构体字段存在。

   ```go
   // 源表结构包含 name age 两个字段
   type User struct {
   	Sex bool
   }
   err := db.Migrator().AddColumn(&User{}, "Sex")
   if err != nil {
   	fmt.Printf("添加字段错误,err:%s\n", err)
   	return
   }
   ```

2. DropColumn 删除字段

   ```go
   err := db.Migrator().DropColumn(&User{}, "Age")
   if err != nil {
   	fmt.Printf("删除字段错误,err:%s\n", err)
   	return
   }
   ```

3. RenameColumn 修改字段名

   `注意`

   - 必须先声明模型。
   - 修改的字段名在对应的数据表必须存在，修改的字段名和修改后的字段名必须定义在结构体内。

   ```go
   type User struct {
   	Name     string
   	UserName string
   }
   err := db.Migrator().RenameColumn(&User{}, "name", "user_name")
   if err != nil {
   	fmt.Printf("修改字段名错误,err:%s\n", err)
   	return
   }
   ```

4. HasColumn 查询字段是否存在

   ```go
   isExistField := db.Migrator().HasColumn(&User{}, "name")
   fmt.Printf("name字段是否存在:%t\n", isExistField)
   isExistField = db.Migrator().HasColumn(&User{}, "user_name")
   fmt.Printf("user_name:%t\n", isExistField)
   ```

### 数据库表的索引接口

1. CreateIndex 为字段创建索引

   `注意`

   - 必须先声明模型。
   - 必须先在声明模型中使用标签`gorm:index`定义索引。

   ```go
   type User struct {
   	gorm.Model
   	Name string `curd:"size:255;index:idx_name,unique"`
   }
   // 为 Name 字段创建索引,两种方法都可以
   db.Migrator().CreateIndex(&User{}, "Name")
   db.Migrator().CreateIndex(&User{}, "idx_name")
   ```

   

2. DropIndex 为字段删除索引

   ```go
   db.Migrator().DropIndex(&User{}, "Name")
   db.Migrator().DropIndex(&User{}, "idx_name")
   ```

   

3. HasIndex 检查索引是否存在

   ```go
   isExists := db.Migrator().HasIndex(&User{}, "idx_name")
   fmt.Printf("idex_name是否存在:%t\n", isExists)
   
   db.Migrator().CreateIndex(&User{}, "idx_name")
   isExists = db.Migrator().HasIndex(&User{}, "idx_name")
   fmt.Printf("idex_name是否存在:%t\n", isExists)
   ```

   

4. RenameIndex 修改索引名

   `注意`

   - 必须先声明模型。
   - 必须先在声明模型中使用标签`gorm:index`定义索引。

   ```go
   type User struct {
   	gorm.Model
   	Name  string `curd:"size:255;index:idx_name,unique"`
   	Name2 string `curd:"size:255;index:idx_name_2,unique"`
   }
   db.Migrator().RenameIndex(&User{}, "idx_name", "idx_name_2")
   ```

   

以上涉及的所有接口，建议都跑一边看看效果~

## 小结

Gorm的迁移接口功能很丰富，AutoMigrate 就适用于大多数的迁移，如果需要更加个性化的迁移工具 ，GORM 提供的一个通用数据库接口。

```go
// returns `*sql.DB`
db.DB()
```

迁移接口的方法，确实给开发工作带来了方便，但是个人建议除非特殊原因，否则尽量通过`在声明模型中修改数据库表的字段和索引`。