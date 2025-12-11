package main

import (
	"os"
	"syscall"
)

func IgnoreAudioWarnings() {
	devNull, _ := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	syscall.Dup2(int(devNull.Fd()), int(os.Stderr.Fd()))
	devNull.Close()
}
