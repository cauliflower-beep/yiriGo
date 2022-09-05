package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type T_msg_event struct {
	Event_id        int       // 事件id
	Event_class_id  int       // 从属事件分类id
	Event_name      string    // 事件名
	Event_desc      string    // 事件描述
	Template_id     int       // 使用模板的id
	Buss_type       int       // 业务类型
	Push_type       int       // 推送目标 0组播 1广播
	Group_push_type int       // 分组推送类型 0：自定目标（uid） 1：自定目标（资金账号）2：推送给持仓用户 3：推送给自选用户
	Create_user     string    // 创建人
	Create_time     time.Time // 创建时间
	Update_user     string    // 修改人
	Update_time     time.Time // 修改时间
	Status          int       // 状态：1有效 0无效
}

//func (ArticleTag) TableName() string {
//	return "t_msg_event"
//}

func main() {
	dsn := "tafadmin:tafadmin@2022@tcp(192.168.101.161:3306)/db_center_manager?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open mysql err:", err)
	}
	fmt.Println(_db)
	sql := "UPDATE t_msg_event SET group_push_type = 6 WHERE event_id = 13;"
	_db.Exec(sql) // 执行
	fmt.Println("update done")
	var event T_msg_event
	query := "select * from t_msg_event where event_id = 13"
	_db.Raw(query).Scan(&event)
	//fmt.Println(event)
	msg_event, _ := json.Marshal(event)
	fmt.Println(string(msg_event))
}
