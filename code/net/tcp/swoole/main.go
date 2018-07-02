package main

import(
	"fmt"
	"net"
	"flag"
	"os"
	// "strconv"
)

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.String("port", "9503", "port")

func main(){
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan string)

	go handWrite(conn, done)
	go handlRead(conn, done)

	fmt.Println(<-done)
	fmt.Println(<-done)

}

func handlRead(conn net.Conn, done chan string){
	buf := make([]byte, 1024)
	reqLen , err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}

	fmt.Println(string(buf[:reqLen-1]))

	done <- "Read"
}

func handWrite(conn net.Conn, done chan string){
	_, e := conn.Write([]byte("hello , message from go \r\n" ))

	if e != nil {
		fmt.Println("Error of send message becase of ", e.Error())
		return
	}


	// for i := 10; i > 0; i-- {
	// 	fmt.Println(i)
	// 	_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))

	// 	if e != nil {
	// 		fmt.Println("Error of send message becase of ", e.Error())
	// 		break
	// 	}

	// 	if i == 1{
	// 		conn.Write([]byte("\r\n"))
	// 	}
	// }


	done <- "Sent"
}