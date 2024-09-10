package _log

import "testing"

func TestLogDefault(t *testing.T) {
	logDefault()
}

func TestLogWithFlags(t *testing.T) {
	logWithFlags()
}

func TestLogWithPrefix(t *testing.T) {
	logWithPrefix()
}

func TestLog2file(t *testing.T) {
	log2File()
}
