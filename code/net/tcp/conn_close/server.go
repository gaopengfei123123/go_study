package main

import (
	"log"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()

	var buf = make([]byte, 10)
	log.Println("start to read from con")
	n, err := c.Read(buf)

	if err != nil {
		log.Println("conn read error:", err)
	} else {
		log.Printf("read %d bytes, connect is %s \n", string(buf[:n]))
	}

	n, err = c.Write(buf)
	if err != nil {
		log.Println("conn write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s\n ", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}

		log.Println("accept a new connection")
		go handleConn(c)
	}
}
