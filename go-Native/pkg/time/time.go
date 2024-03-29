package main

import (
	"fmt"
	"strconv"
	"time"
)

const TIME_FORMAT = "2006-01-02"

func main() {
	now := time.Now()
	fmt.Println(now.Unix())
	// 24小时之内的计算
	//t1, _ := time.ParseDuration("12h")
	//m1 := now.Add(t1).Unix()
	//m2 := uint32(m1)
	//fmt.Println(m1, m2)
	// 减5分钟
	t2, _ := time.ParseDuration("-5m")
	m3 := now.Add(t2)
	fmt.Println(m3.Unix())

	// 本地时间转 UTC 11.05
	date := time.Now().UTC().Format(TIME_FORMAT)
	date2 := time.Unix(1666908000, 0).UTC().Format(TIME_FORMAT) // 东8区会比标准时间 UTC 快8个小时
	fmt.Println(date, date2)

	// 当前时间戳转字符串 12.08
	fmt.Println("当前时间戳：", strconv.FormatInt(now.Unix(), 10))

	// 前一天的日期字符串
	lastDate := time.Now().AddDate(0, 0, -1).Format("20060102")
	fmt.Println(lastDate)

}
