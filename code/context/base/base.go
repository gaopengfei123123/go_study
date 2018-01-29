package main

import (
	"context"
	"fmt"
)

func main() {
	// 这里在一个函数里使用非缓冲 channel 自收自发
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n:= 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	// 当 goroutine 结束时关闭,防止 goroutinue 泄露
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n >= 5 {
			break
		}
	}
}