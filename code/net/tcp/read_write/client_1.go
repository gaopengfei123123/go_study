package main

import(
	"log"
	"net"
	"time"
)

func main(){
	log.Println("start dial...")
	conn, err := net.Dial("tcp",":8899")
	if err != nil {
		log.Println("error: " , err)
		return
	}

	defer conn.Close()
	log.Println("dial ojbk")
	time.Sleep(time.Second * 1000)
}