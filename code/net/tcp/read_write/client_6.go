package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		log.Println("dial error:", err)
		return
	}

	defer conn.Close()
	log.Println("dial ok")

	data := make([]byte, 65536)
	var total int

	// 模拟写超时, 同样请求 server_5.go 程序, 不过对写入动作进行了一个超时判定
	for {
		conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))
		n, err := conn.Write(data)
		if err != nil {
			total += n
			log.Printf("write %d bytes, error: %s", n, err)
			break
		}
		total += n
		log.Printf("write %d bytes this time, %d bytes in total \n", n, total)
	}

	log.Printf("write %d bytes in total\n", total)
	time.Sleep(time.Second * 10000)
}
