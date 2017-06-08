package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	tm := time.Unix(timestamp, 0)
	timer := tm.Format("2006-01-02 03:04:05 PM")
	fmt.Println(timer)
}
