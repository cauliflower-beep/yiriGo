package _database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func queryRowDemo() {
	// 查询某条记录
	querySql := "select id,name,birthday from t_player where id = ?"
	row := db.QueryRow(querySql, 1)
	var p = struct {
		id       int
		name     string
		birthday string
	}{}
	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	if err = row.Scan(&p.id, &p.name, &p.birthday); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("玩家id %d 不存在\n", 1)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("玩家id:%d, name:%s, birthday:%s\n", p.id, p.name, p.birthday)
	}
}
