package main
import (
	"log"
	"net"
	"time"
)

// 这个和 client_1 时的问题就凸现出来, 当 server 的 backlog 到达上限后则在之后的链接全部都阻塞住了
func main() {
	var sl []net.Conn

	for i:=1; i<1000; i++{
		conn := extablishConn(i)

		if conn != nil {
			sl = append(sl, conn)
		}
	}

	time.Sleep(time.Second * 1000)
}

func extablishConn(i int) net.Conn{
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		log.Printf("%d: dial error: %s", i, err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}