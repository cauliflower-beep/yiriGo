package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

//cron 启动的定时任务，如果此次周期内没有处理完会怎么样？
// 每次任务触发，都是开了一个协程处理，所以下次任务触发不影响此次任务执行

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	sb := strings.Builder{}
	bigInt := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		randInt, _ := rand.Int(rand.Reader, bigInt)
		sb.WriteRune(rune(charset[randInt.Int64()]))
	}
	return sb.String()
}

func main() {
	c := cron.New()

	// 让这个定时任务每1分钟启动一次
	_ = c.AddFunc("0 0/1 * * * ?", func() {
		// 每次启动都随机生成一个字符串
		randomString := generateRandomString(4)
		for i := 0; i < 10; i++ {
			// 每 10s 打印一下随机字符串
			fmt.Printf("current string|%s\n", randomString)
			time.Sleep(10 * time.Second)
		}
	})
	c.Start()

	// 保证主go程不退出
	select {}
}
