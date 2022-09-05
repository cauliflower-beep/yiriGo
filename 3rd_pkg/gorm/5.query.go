package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

//type T_task_pool struct {
//	Task_id string `gorm:"column:task_id"`
//	//Push_msg_id    int64
//	Event_id       int       `gorm:"column:event_id"`
//	Push_req       string    `gorm:"column:push_req"`
//	Access_id      string    `gorm:"column:access_id"`
//	Status         int       `gorm:"column:status"`
//	Update_time    time.Time `gorm:"column:update_time"`
//	Create_time    time.Time `gorm:"column:create_time"`
//	Push_exec_time time.Time `gorm:"column:push_exec_time"`
//	Push_target    string    `gorm:"column:push_target"`
//}

type T_task_pool struct {
	Task_id  string
	Event_id int
	//Push_req       string
	Access_id      string
	Status         int
	Update_time    time.Time
	Create_time    time.Time
	Push_exec_time time.Time
	Push_target    string
}

func main() {
	dsn := "upchina:upchina2017@tcp(172.16.8.45:3306)/db_pmsg_hfzq?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open mysql err:", err)
	}

	sql := "select * from t_task_pool where create_time > \"2022-07-25 10:33:15\" order by task_id asc;"
	//sql := "select task_id,event_id,push_req,access_id,`status`," +
	//	"update_time,create_time,push_exec_time,push_target " +
	//	"FROM t_task_pool where create_time > \"2022-07-25 10:33:15\" " +
	//	"order by task_id asc;"
	//sql := "select task_id,event_id,access_id,`status`," +
	//	"update_time,create_time,push_exec_time,push_target " +
	//	"FROM t_task_pool where create_time > \"2022-07-25 10:33:15\" " +
	//	"order by task_id asc;"
	var pool []T_task_pool
	_db.Raw(sql).Scan(&pool)
	fmt.Println("{querySql res}|%v", pool)

	var tasks []string
	for _, task := range pool {
		t, _ := json.Marshal(task)
		tasks = append(tasks, string(t))
		//fmt.Println(tasks)
	}
	// 拼接返回
	taskRecord := strings.Join(tasks, "\n")
	fmt.Println(taskRecord)
	//fmt.Println("{querySql res}|%s", taskRecord)
}
