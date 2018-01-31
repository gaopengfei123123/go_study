package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bufio"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the url...")
	url, err := inputReader.ReadString('\n')
	url = strings.Trim(url, "\n")
	fmt.Println("url:"+ url)
	checkError(err)

	res, err := http.Get(url)
	checkError(err)

	data, err := ioutil.ReadAll(res.Body)
	checkError(err)

	fmt.Printf("Got: %q",string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}