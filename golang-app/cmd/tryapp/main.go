package main

import (
	"golang-app/internal/logger"
	"runtime/debug"
)

// Start of Selection
func main() {

	log := logger.NewLogger()

	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			log.Error("panicが発生しました: ", "error", err, "stack", stack)
		}
	}()

	if err := someFunc(); err != nil {
		log.Error("error", "error", err)
		return
	} else {
		log.Info("success")
	}
}

func someFunc() error {
	panic("パニック")
}
