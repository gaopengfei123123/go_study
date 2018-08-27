package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client_2.go hello")
		return
	}

	log.Println("start dial...")
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		log.Println("error: ", err)
		return
	}

	defer conn.Close()
	data := os.Args[1]
	conn.Write([]byte(data))
	log.Println("dial ojbk")
}
