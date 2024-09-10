package _logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

/*
logrus是一个强大且灵活的结构化日志库，提供了丰富的功能和可定制性
它支持多种输出格式、日志级别、钩子和格式器
*/

func base() {
	log := logrus.New() // 创建一个日志记录器

	log.Println("base msg")

	log.WithFields(logrus.Fields{
		"key": "val", // 添加结构化日志字段
	}).Info("base msg with fields.")
}

/*
主要功能和特点：
1.结构化日志记录：logrus支持使用字段来记录结构化日志，使得日志更加易于查询和过滤；
2.多种输出格式：logrus支持多种输出格式，如Json、文本等，方便根据需求选择合适的输出格式；
3.日志级别：logrus支持多个日志级别：debug、info、warn、error等，可以根据需要灵活选择；
4.钩子和格式器：logrus提供了钩子和格式器的机制，可以扩展其功能和定制日志的输出行为。
*/

/*
自定义日志输出格式
*/
func jsonFormat() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{}) // 使用Json格式输出

	log.WithFields(logrus.Fields{
		"key": "val",
	}).Info("json format msg with fields.")
}

/*
添加钩子(hooks)：
logrus支持在日志记录期间触发钩子来执行额外的逻辑
钩子可以用于发送日志到外部服务、记录日志到数据库等。
*/
func logWithHooks() {
	log := logrus.New()
	log.AddHook(&myHook{})
	log.Info("log msg with hooks")
}

// 自定义钩子
type myHook struct{}

// Levels 用于指定钩子监听的日志级别
func (h *myHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *myHook) Fire(e *logrus.Entry) error {
	fmt.Println("111")
	// 用于在日志记录时触发钩子逻辑，可以在这里执行额外的日志处理操作
	return nil
}

// https://blog.csdn.net/qq_42531954/article/details/136491931
