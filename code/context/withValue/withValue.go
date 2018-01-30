package main

import (
	"context"
	"fmt"
)

type ctxValue struct {
	Name string
	Sex string
}
func main() {
	type favContextKey string 
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return 
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctxVal := ctxValue{"gpf","male"}

	// 添加一个 key-value 键值对,所有使用该 context 的都能访问到
	ctx := context.WithValue(context.Background(), k, ctxVal)

	f(ctx, k)
	f(ctx, favContextKey("color"))
}