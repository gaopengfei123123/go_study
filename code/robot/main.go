package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("An Error occurred:%s\n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Hello %s! What can I do for you?\n", name)
	}

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An Error occurred:%s\n", err)
			continue
		}

		input = input[:len(input)-1]
		input = strings.ToLower(input)

		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didn't catch you")

		}
	}

}
