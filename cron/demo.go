package main

import (
	"github.com/robfig/cron"
	"log"
)

/*
企业项目完整之后，经常会有一些定时任务，例如备份、检查数据等，需要经常定时跑，如果我们能把这些定时任务，做成一个业务系统，通过与liunx下cron一样的规则，来表示定时执行的规则，岂不美哉。

Golang的cron包帮忙解决这个问题。
*/
func main() {
	log.Println("Starting...")

	c := cron.New()
	/*
		1.cron表达式共有六个域，分别表示：Seconds,Minutes,Hours,DayofMonth,Month,DayofWeek

		2.每个域的基本格式如下：
			字段名				是否必须			允许的值				允许的特定字符
			秒(Seconds)			是				0-59				* / , –
			分(Minutes)			是				0-59				* / , –
			时(Hours)			是				0-23				* / , –
			日((Day of month)	是				1-31				* / , – ?
			月((Month)			是				1-12 or JAN-DEC		* / , –
			星期(Day of week)	否				0-6 or SUM-SAT		* / , – ?

			1）月(Month)和星期(Day of week)字段的值不区分大小写，如：SUN、Sun和 sun是一样的。
			2）星期(Day of week)字段如果没提供，相当于是 *

		3.特殊字符：
			1）星号(*)：表示cron表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月；
			2）斜线(/)：表示增长间隔，如第1个字段(minutes)值是3-59/15，表示每小时的第3分钟开始执行一次，
					   之后每隔 15 分钟执行一次（即 3、18、33、48这些时间点执行），这里也可以表示为：3/15；例如：spec := "/5 " 每隔5s执行一次；
			3）逗号(,)：用于枚举值，如第6个字段值是 MON,WED,FRI，表示星期一、三、五执行；例如: spec := " 52,54 9 " 每天9:52分和9:54分的每秒都执行一次；
			4）连字号(-)：表示一个范围，如第3个字段的值为 9-17 表示 9am到 5pm直接每个小时（包括9和17）；
						例如：spec := "15-30 *" //每分钟的15-30s执行定时任务；
			5）问号(?)：只用于日(Day of month)和星期(Day of week)，表示不指定值，可以用于代替 *
	*/
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
	})

	c.Start()

	//t1 := time.NewTimer(time.Second * 10)
	//for {
	//  select {
	//  case <-t1.C:
	//      t1.Reset(time.Second * 10)
	//  }
	//}
	select {}
}
