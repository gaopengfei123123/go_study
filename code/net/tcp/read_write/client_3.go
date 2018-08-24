package main

import(
	"log"
	"net"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client2.go YOUR_CONTENT")
		return
	}

	log.Println("start dial...")
	conn, err := net.Dial("tcp",":8899")
	if err != nil {
		log.Println("error: " , err)
		return
	}

	defer conn.Close()
	data := os.Args[1]
	conn.Write([]byte(data))
	log.Println("dial ojbk")
}