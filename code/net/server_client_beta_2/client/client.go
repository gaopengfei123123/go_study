package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var conn net.Conn
	var error error 
	var inputReader *bufio.Reader
	var input string
	var clientName string

	conn, error = net.Dial("tcp","localhost:50000")
	checkError(error)

	// 获取链接用户名
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("First, what's your name?")
	clientName, _ = inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to shutdown server.")
		input, _ =  inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input,"\n")
		// Q是直接退出
		if trimmedInput == "Q" {
			return 
		} 

		
		_, error = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		checkError(error)
		// 输入 SH 将 server 端和 client 都关闭
		if trimmedInput == "SH" {
			return 
		} 
	}

}

func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error()) // terminate program
	}
}