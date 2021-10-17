package main

import(
	"context"
	"fmt"

	diandiansignal "github.com/diandianl/p2p-proxy/signal"
)

func main() {
	ctx := context.Background()

	done, ctx := diandiansignal.SetupInterruptHandler(ctx)

	defer done()

	fmt.Println("hello Starry Proxy...")
}
