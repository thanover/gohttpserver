package main

import (
	"fmt"
	"net"
	"os"

	"github.com/thanover/gohttpserver/utils"
	logger "github.com/thanover/gologger"
)

var AppLogger logger.Logger

func main() {
	logger.Initialize("./logs")
	AppLogger = logger.GetLogger()

	// set up a listener on a TCP port (5544)
	ln, err := net.Listen("tcp", ":5544")
	if err != nil {
		AppLogger.Error(err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	AppLogger.Info("Listening on :5544")

	for {
		conn, err := ln.Accept()
		if err != nil {
			AppLogger.Error(err.Error())
			continue
		}
		go handleConnection(conn)
	}
	//receive the message and parse it

	//handle the request

	//respond with proper message

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	AppLogger.Info(fmt.Sprintf("Connection from %s\n", conn.RemoteAddr()))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		AppLogger.Info(fmt.Sprintf("Received:\n \n %s", buf[:n]))
		AppLogger.Info(fmt.Sprintf("%d", buf[:n]))

		stringArray := utils.BytesToStrings(buf[:n])
		AppLogger.Info(fmt.Sprint(stringArray))

		conn.Write(buf[:n])
	}
}
