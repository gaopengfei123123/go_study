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
		if trimmedInput == "Q" {
			return 
		} 

		_, error = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		checkError(error)
	}

}

func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error()) // terminate program
	}
}