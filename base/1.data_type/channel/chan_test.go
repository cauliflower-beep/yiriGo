package main

import "testing"

func TestBlockChanNoBuffer(t *testing.T) {
	blockChanNoBuffer()
}

func TestUnBlockChanNoBuffer(t *testing.T) {
	unblockChanNoBuffer()
}

func TestBlockChanBuffer(t *testing.T) {
	blockBufferChan1()

	// blockBufferChan2()
}

func TestUnBlockChanBuffer(t *testing.T) {
	// unblockBufferChan1()

	unblockBufferChan2()
}
