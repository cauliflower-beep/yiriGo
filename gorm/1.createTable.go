package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	Name     string
	Age      int
	Brithday time.Time
}

func main() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open mysql err:", err)
	}
	user := User{Name: "lufy", Age: 22, Brithday: time.Now()}
	_ = db.Create(&user)
}
