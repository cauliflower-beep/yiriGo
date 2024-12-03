package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func main() {
	logger := &lumberjack.Logger{
		Filename:   "test.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		LocalTime:  true,
	}
	defer logger.Close()

	log.SetOutput(logger)

	for {
		log.Println("test111")
	}
}
