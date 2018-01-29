package main
import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(4 * time.Second)
	// d := time.Unix(86400, 0)

	// 设置超时时间, 这里是基于绝对时间的倒计时,如果晚于执行时的时间会立即执行
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	// 这里使用 time.After 来模拟其他的 goroutinue 操作,一旦到达时间将会中断执行
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
/**
输出:
overslept
overslept
overslept
context deadline exceeded
*/