package utils

import (
	"log"
	"os"
	"runtime/debug"
)

var Logger = log.New(os.Stdout, "[JetChatClientGo] ", log.LstdFlags)

func LogError(err error) {
	debug.PrintStack()
	Logger.Fatal(err)
}
