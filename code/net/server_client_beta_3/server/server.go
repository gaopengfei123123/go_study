package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	// 通过启动参数设置服务器地址和接口
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s",flag.Arg(0),flag.Arg(1))
	// fmt.Println("current host", hostAndPort)
	listener := initServer(hostAndPort)

	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func connectionHandler(conn net.Conn) {
	//获取client 端 ip 信息
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ",connFrom)
	sayHello(conn)

	for {
		// 承接最大字符量的切片,每25个字符为一段
		ibuf := make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0

		// 根据错误的返回类型来进行不同的处理
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")

}

// 将字节流打印出来
func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func sayHello(to net.Conn){
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote " + string(wrote) + " bytes.")
}

func initServer(hostAndPort string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed:'" + hostAndPort + "'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTcp:")
	println("Listening to: ", listener.Addr().String())
	return listener
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}
