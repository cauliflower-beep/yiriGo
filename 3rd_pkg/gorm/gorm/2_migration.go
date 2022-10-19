package gorm

/*
	迁移也就是对表进行操作
*/

func init() {
	/*
		当 model 存在自定义字段且未实现
			Scan(value interface{}) error
			Value() (driver.Value, error)
		两个接口时，直接创建表会失败.

		AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。
		AutoMigrate 会创建表\缺少的外键\约束\列和索引，并且会更改现有列的类型（如果其大小、精度、是否为空可更改）。
		但 不会 删除未使用的列，以保护您的数据。
	*/
	//DB.Migrator().DropTable(Character{}) // 删表重建
	//_ = DB.Migrator().AutoMigrate(Character{})
}
