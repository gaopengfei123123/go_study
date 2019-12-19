package main

import (
	"log"
	"time"
)

func main() {
	test()
	bigSlowOperation("test input")
}

func bigSlowOperation(input interface{}) (output interface{}) {
	defer trace("bigSlowOperation", input)()
	// ...lots of work…
	time.Sleep(2 * time.Second) // simulate slow
	output = "result"
	log.Println("done")
	return
}
func trace(funcName string, args ...interface{}) func() {
	start := time.Now()
	log.Printf("enter %s", funcName)
	return func() {
		log.Printf("func args: %s \n", args)
		log.Printf("exit %s (%s)", funcName, time.Since(start))
	}
}

func test() {
	for i := 0; i < 5; i++ {
		defer func(index int) {
			log.Printf("i: %d, index: %d \n", i, index)
		}(i)
	}
}
