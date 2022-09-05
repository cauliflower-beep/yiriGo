package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ArticleTag struct {
	article_id  int64
	tag_id      int64
	status      int64
	create_time int64
}

func (ArticleTag) TableName() string {
	return "t_article_tag"
}

func main() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open mysql err:", err)
	}
	tag := ArticleTag{article_id: 1, tag_id: 0, status: 0, create_time: 0}
	_ = db.Create(&tag)
}
