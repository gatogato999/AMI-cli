package main

import (
	"fmt"
	"log"
	"net"

	"github.com/staskobzar/goami2"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5038")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goami2.NewClient(conn, "admin", "passpass")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	for msg := range client.AllMessages() {
		fmt.Println(msg.String())
	}
}
