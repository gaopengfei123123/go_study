package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp","localhost:5000")
	if err != nil {
		fmt.Println("Error dialing",err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First,what's youre name?")
	clientName, _ := inputReader.ReadString('\n')
	srimmedClient := strings.Trim(clientName,"\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input,"\n")
		if trimmedInput == "Q" {
			return
		}

		_, err := conn.Write([]byte(srimmedClient + " says: " + trimmedInput))
		if err != nil {
			fmt.Println("Error message",err.Error())
			return
		}
	}
}