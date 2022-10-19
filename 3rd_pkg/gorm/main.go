package main

import (
	"yiriGo/3rd_pkg/gorm/gorm"
	_ "yiriGo/3rd_pkg/gorm/gorm"
)

func main() {
	// 创建记录
	//gorm.CreateRec()
	//gorm.SelectCreate()
	//gorm.OmitCreate()
	//gorm.BatchCreate()

	// 查询记录
	gorm.GetOnceRec()
}
