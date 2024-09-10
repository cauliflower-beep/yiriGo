package _logrus

import "testing"

func TestBase(t *testing.T) {
	base()
	/*
		如果出现这个报错：
		runtime: bad pointer in frame github.com/konsorten/go-windows-terminal-sequences.EnableVirtualTerminalProcessing at 0xc0000294c8: 0x784
		fatal error: invalid pointer found on stack
		检查下这个依赖包的版本：
		github.com/konsorten/go-windows-terminal-sequences
		如果是v1.0.1,需要升级到v1.0.3以上
	*/
}

func TestJsonFormatter(t *testing.T) {
	jsonFormat()
}

func TestLogWithHook(t *testing.T) {
	logWithHooks()
}
