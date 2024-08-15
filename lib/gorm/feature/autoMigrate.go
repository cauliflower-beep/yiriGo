package main

import (
	"gorm.io/gorm"
	conn "yiriGo/lib/gorm/Connection"
)

/*
	AutoMigrate方法可以自动同步Go结构体到数据库表，并确保这些表的结构与Go模型匹配

	AutoMigrate不会删除已存在的列，即使Go结构体中不再包含这些列。
	如果需要删除列，需要手动执行Sql语句，或者先删除表再重新创建
*/

type books struct {
	Title  string  `gorm:"column:title"`
	Price  float64 `gorm:"column:price"`
	Author string  `gorm:"column:author"`
	gorm.Model
}

func (b *books) TableName() string {
	return "t_books"
}

func CreateTableByAutoMigrate() {
	_ = conn.DB.Exec("DROP TABLE IF EXISTS `t_books`;") // 先删表
	_ = conn.DB.AutoMigrate(&books{})                   // 自动创建表
}

func AddFieldByAutoMigrate() {
	_ = conn.DB.AutoMigrate(&books{}) // 利用AutoMigrate扩展数据库字段
}
