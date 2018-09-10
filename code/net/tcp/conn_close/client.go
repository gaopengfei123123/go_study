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
		log.Println("dial err:", err)
		return
	}

	conn.Close()
	log.Println("close ok")

	var buf = make([]byte, 32)
	n, err := conn.Read(buf)

	if err != nil {
		log.Println("read error:", err)
	} else {
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}

	n, err = conn.Write(buf)
	if err != nil {
		log.Println("write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s \n", n, string(buf[:n]))
	}

	time.Sleep(time.Second * 1000)

}
