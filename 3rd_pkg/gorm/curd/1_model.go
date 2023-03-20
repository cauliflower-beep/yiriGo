package curd

import (
	"database/sql"
	"gorm.io/gorm"
)

/*
	curd 允许使用自定义类型，但必须实现两个接口：
		Scan(value interface{}) error
		Value() (driver.Value, error)

	这样 curd 才能正常识别转化，否则会读不出来：
		define a valid foreign key for relations or implement the Valuer/Scanner interface

	因为无论你怎么约定，例如下面的自定义 Hobby 字段，数据库根本不存在这样一个数组类型的切片
	curd 认为你的自定义类型，要满足两个条件：
		1.首先要有一个能够获取 value 的方法；
		2.其次需要知道，scan 之后，映射到你自定义类型的哪个字段上
*/
type Hobby []string
type Job struct {
	Title    string
	Location string
}
type Character struct {
	/*
		内嵌 curd 预定义的结构体，具体字段看结构体内容
		其中的 CreatedAt\UpdateAt 是 curd 自动追踪填充的
	*/
	gorm.Model
	/*
		colunm 是很常用的一个 curd 标签
		很多时候，数据库的表并不是你创建的，这个时候出现字段匹配不上的现象很正常，
		可以使用 column 来做一个匹配
	*/
	Name     string `curd:"column:name"`
	Age      uint8  `curd:"check:age > 12"` // 为 age 字段创建检查约束
	Fight    uint8
	Birthday int64 `curd:"serializer:unixtime;type:time"` // 跟 int64 类型可以互转
	Pal      string
	/*
		sql自定义的一种类型
		读数据库例如 rows.Scan 的时候，有时会返回空值。
		如果直接把空值复制给 string 变量，会出错。
		这时就需要 sql.NullString 类型的变量了，他的结构体内容自查。
		第一个存字符串，第二个存的是是否为 Null
		把变量声明成 sql.NullString 类型，再加上一个判断，就可以避免上面的问题
	*/
	Enemy    sql.NullString
	Hobby    Hobby `curd:"serializer:json"` // 指定字段如何在数据库中进行序列化和反序列化
	JobInfo1 Job   `curd:"embedded;embeddedPrefix:job_"`
	JobInfo2 Job   `curd:"type:bytes;serializer:gob"`
}
