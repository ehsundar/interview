package main

import (
	"context"
	"fmt"

	"github.com/ehsundar/interview.git/consistenthashing/internal/resolver"
)

func main() {
	ctx := context.Background()
	key := "amir mohammad"

	r := resolver.NewResolver([]string{
		"server_0",
		"server_1",
		"server_2",
		"server_3",
	}, resolver.WithVirtualizationFactor(1))
	target, _ := r.Resolve(ctx, key)

	fmt.Printf("key: %s, target: %s\n", key, target)

	r.PrintConfiguration()

	r.AddTarget("server_100")
	r.AddTarget("server_101")
	r.PrintConfiguration()
}
