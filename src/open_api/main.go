package main

import(
	"fmt"
	"encoding/json"
	"net"
	"flag"
	"os"
	// "strconv"
)

type requestTmp struct{
	Src string 		`json:"src"`
	Type string		`json:"type"`
	IP string		`json:"ip"`
	Route string	`json:"route"`
	LogID string 	`json:"log_id"`
}

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.String("port", "6666", "port")

func main(){
	flag.Parse()
	fmt.Println("hello", string(*host), string(*port))

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
	params := requestTmp{
		"hers",
		"h5",
		"127.0.0.1",
		"auth/user/register",
		"5b39c4f8be099",
	}
	b, err := json.Marshal(params)
	checkError(err)

	_, e := conn.Write(b)

	if e != nil {
		fmt.Println("Error of send message becase of ", e.Error())
		return
	}

	done <- "Sent"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}