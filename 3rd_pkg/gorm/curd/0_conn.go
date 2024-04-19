package curd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB
var err error

var dsn = "root:admin123@tcp(127.0.0.1:3306)/test111?charset=utf8mb4&parseTime=True&loc=Local"

/*
推荐使用增加配置的方式链接 mysql
因为比如字符类型，gorm在创建表的时候会默认使用比较大的 text 类型，而我们日常使用最多的其实是 varchar
*/
func init() {
	DB, err = gorm.Open(mysql.New(mysql.Config{
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
}

// init 已经生成了一个数据库连接池，所以我们只需要 set 即可
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB() // 实际上操作数据的是这个 sqlDB 它本质上是一个连接池
	if err != nil {
		log.Println(err)
		return
	}
	/*
		下述属性如果不知道怎么设置，可以使用默认的
	*/
	sqlDB.SetMaxIdleConns(5)  // 最大空闲数
	sqlDB.SetMaxOpenConns(10) // 最多同时存在的连接数
	/*
		最大存活时间
		一旦某个链接达到这个时间，不管它是不是空闲的，都会被杀掉重建
		这样其实是防止某个链接不断被使用、释放的过程中出现意外
		或者某个链接有bug，每次使用都会导致内存的累加，释放之后可以使内存复原
	*/
	sqlDB.SetConnMaxLifetime(time.Hour)
}
