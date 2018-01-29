package main

import(
	"context"
	"fmt"
	"time"
)

func main () {
	// withTiemOut 就是 withDeadLine 的一层封装, 返回的是 WithDeadline(parent, time.Now().Add(timeout)),基于当前时间的倒计时
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	LOOP:
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break LOOP
		}
	}
	
}