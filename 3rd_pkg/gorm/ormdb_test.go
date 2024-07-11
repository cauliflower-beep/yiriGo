package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
)

func TestDbConn(t *testing.T) {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/ypwork?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,  // DSN data source name
		DefaultStringSize:        256,  // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:   true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//Conn: 	// 如果已有现存的链接，则无需指定 DSN 链接字符串，只需要指定一个现有的连接即可
	}), &gorm.Config{
		// 为了方便学习观察，看看 curd 将我们对象关系映射到数据库的sql语句长什么样，可以通过日志打印一下
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别
	})
	if err != nil {
		log.Println(err)
		return
	}

	var recs []map[string]interface{}
	DB.Table("subscription_column").Find(&recs)
	fmt.Println(recs)

}
