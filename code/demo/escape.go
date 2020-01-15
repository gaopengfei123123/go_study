package main

import "fmt"

// go:noinline
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// go:noinline
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

// i := 0

// // b := 1
// // fmt.Printf("%#+v \n", &b)

// nextInt := func() int {
// 	i++
// 	fmt.Printf("%#+v \n", &i)
// 	return i
// }
