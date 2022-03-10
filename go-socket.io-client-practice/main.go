package main

import (
	"fmt"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(transport.GetDefaultWebsocketTransport())

	c, err := gosocketio.Dial(
		"wss://api.vtbs.moe/socket.io",
		transport.GetDefaultWebsocketTransport())

	if err != nil {
		log.Fatal("err:", err)
	}
	err = c.On("info", func(args ...interface{}) {
		fmt.Println(args)
	})
	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Println("Disconnected")
	})
	time.Sleep(10 * time.Second)
}
