package main

import (
	"demo-test/internal/log"
)

func main() {
	log.LoggerInit()
	log.Logger.Info("Hello, World!")
}
