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

	// 模拟写阻塞: 当双方的OS 协议栈中的数据缓冲占满时就会出现阻塞, 当 server 端输出一部分之后, client 就又能往里面塞数据了
	for {
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
