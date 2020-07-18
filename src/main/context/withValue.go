package main

import (
	"context"
	"fmt"
)

type favContextKey string

func main() {

	f := func(ctx context.Context, key favContextKey) {
		if v := ctx.Value(key) ; v != nil {
			fmt.Println("found value : " , v)
			return
		}

		fmt.Println("not found key: ", key)
	}

	key := favContextKey("hell")

	//返回一个 带有 value 的ctx
	ctx := context.WithValue(context.Background(), key, "golang")

	//检查key 在否
	f(ctx, key)

	f(ctx, favContextKey("wrwr"))
}
