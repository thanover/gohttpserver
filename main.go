package main

import (
	logger "github.com/thanover/gologger"
)

func main() {
	logger := logger.MakeLogger("./log")
	logger.Info("Test Message")
}
