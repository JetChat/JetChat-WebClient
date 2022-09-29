package utils

import (
	"log"
	"os"
	"runtime/debug"
)

var Logger = log.New(os.Stdout, "[JetChatClientGo] ", log.LstdFlags)

func FatalError(err error) {
	debug.PrintStack()
	Logger.Fatal(err)
}

func LogError(err error) {
	debug.PrintStack()
	Logger.Println(err)
}
