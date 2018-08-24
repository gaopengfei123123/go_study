package main
import (
	"log"
	"net"
	"time"
)

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
	conn, err := net.DialTimeout("tcp", ":8899", time.Second)
	if err != nil {
		log.Printf("%d: dial error: %s", i, err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}