package main
import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	var listener net.Listener
	var error error
	var conn net.Conn
	mapUsers = make(map[string]int)

	fmt.Println("Starting the server ...")

	listener, error = net.Listen("tcp", "localhost:50000")
	checkError(error)

	for {
		conn, error = listener.Accept()
		checkError(error)
		go doServerStuff(conn)
	}
	
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var error error

	// 类似 try catch 的结构
	defer func() {
		fmt.Println("recover painc")
		err := recover();
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)

		if error != nil {
			panic(error.Error())
		}

		checkError(error)
		input := string(buf)

		// 通过远程命令关闭监听进程
		if strings.Contains(input, ": quit") {
			panic("Client shutting down.")
		}

		// 通过远程命令关闭监听进程
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}

		// 关键测 WHO 来列出用户列表
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}

		nameIdx := strings.Index(input,"says")
		sayName := input[0:nameIdx]
		mapUsers[string(sayName)] = 1
		fmt.Printf("Received data: ---%v--- \n", string(buf))
	}
}

func checkError(error error) {
	if error != nil {
		fmt.Println("Error: " + error.Error()) // terminate program
	}
}

// DisplayList 显示当前链接数量
func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}