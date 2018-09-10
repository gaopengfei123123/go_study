package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Println("listen error: ", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			return
		}
		log.Println("accept a new connection")
		go readHandler(c)
	}
}

func readHandler(c net.Conn) {
	defer c.Close()

	// 这里将统一输出内容, 因为一次性不能把连接中的数据都读取出来
	var result = make([]byte, 100)

	for {
		var buf = make([]byte, 10)
		n, err := c.Read(buf)
		log.Printf("read from conn, len: %d \n", n)
		if err != nil && err.Error() == "EOF" {
			log.Println("conn read end:", err)
			break
		}
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		result = append(result, buf...)
	}
	log.Printf("read bytes, content is %s\n", string(result))
}
