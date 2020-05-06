package main

import (
	"context"
	"fmt"

	"github.com/danielsan22/golang-architecture/ex/req"
)

func main() {

	ctx := context.Background()
	ctx = req.WithUserId(ctx, 1)
	ctx = req.WithUserName(ctx, "Daniel")

	name := req.UserName(ctx)
	fmt.Println(*name)

	if id := req.UserId(ctx); id != nil {
		fmt.Println(*id)
	} else {
		fmt.Println("no value")
	}

}
