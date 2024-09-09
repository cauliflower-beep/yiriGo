package _log

import (
	"log"
	"log/slog"
	"os"
)

/*
旧版本log库没有日志分级，且不支持结构化日志输出
slog是1.21引入的新版本日志库，目的就是将结构化日志引入标准库
*/

func defaultSlog() {
	/*
		新旧两个日志包都公开了一个默认的Logger
		区别在于slog包包含了日志级别
	*/
	log.Println("info msg")

	slog.Info("info msg")
}

/*
通过slog.New()方法可以创建自定义logger实例，以实现结构化的日志输出
slog内置了两个结构化处理程序：TextHandler 和 JSONHandler
*/
func jsonHandle() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// 默认都是Info级别，所以Debug条目会被屏蔽
	logger.Debug("debug msg")
	logger.Info("info msg")
	logger.Warn("warn msg")
	logger.Error("error msg")
}

func txtHandle() {
	opt := slog.HandlerOptions{
		Level: slog.LevelDebug, // 设置日志级别
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opt))
	logger.Debug("debug msg")
	logger.Info("info msg")
	logger.Warn("warn msg")
	logger.Error("error msg")
}
