package __update

import (
	"fmt"
	conn "yiriGo/lib/gorm/0_conn"
)

/*
	更新操作 和 创建操作差不多，可直接读文档
	也是分为了更新单列、更新多列、更新选中字段、忽略某些字段进行更新
*/

func UpdData() {
	//sql := "update gormtest set age = 24 where id = 8;" // 更新一条不存在的记录，不会报错，RowsAffected 为8

	sql := "update gormtest set age = 25 where id = 9 and name = '秀才';"

	res := conn.DB.Exec(sql)
	if err := res.Error; err != nil {
		fmt.Printf("[upd exec failed]|%+v\n", err)
	}
	fmt.Println(res.RowsAffected)
}
