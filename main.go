package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/staskobzar/goami2"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5038")
	if err != nil {
		log.Fatal(err)
	}

	// as in /etc/asterisk/manager.d/go.conf  and /etc/asterisk/manager.conf
	client, err := goami2.NewClient(conn, "admin", "passpass")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	file, err := os.Create("asterisk-events.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for msg := range client.AllMessages() {
		fmt.Println(msg.String())

		_, writeErr := fmt.Fprintln(file, msg.String())
		if writeErr != nil {
			log.Println(writeErr)
		}
	}
}
