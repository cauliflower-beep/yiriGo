package main

import (
	"fmt"
	"yiriGo/3rd_pkg/gorm/curd"
	_ "yiriGo/3rd_pkg/gorm/curd"
)

func main() {
	// 创建记录
	//curd.CreateRec()
	//curd.SelectCreate()
	//curd.OmitCreate()
	//curd.BatchCreate()

	// 查询记录
	//curd.GetOnceRec()
	//curd.GetRecs()

	insertSql := "INSERT INTO `test111`.`gormtest` (`id`, `name`, `modify_time`, `nickname`, `age`, `id_card`) VALUES (13, '湘玉', NULL, '掌柜的', 27, '103');"
	res := curd.DB.Exec(insertSql)
	if res.Error != nil {
		fmt.Println("sql exec failed.", res.Error.Error())
		return
	}
	fmt.Println("sql exec succ.", res.RowsAffected)
}
