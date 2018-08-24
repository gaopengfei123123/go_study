
package main

import (
	"log"
	"net"
)

// 请求 server.go 的时候会立即返回, 然而 server 那却发生阻塞
func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
}