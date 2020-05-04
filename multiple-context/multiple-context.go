package multiple

import (
	"context"
	"fmt"
	"time"
)

func MultipleContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "node01")
	go worker(ctx, "node02")
	go worker(ctx, "node03")
	go worker(ctx, "node04")
	time.Sleep(5 * time.Second)
	fmt.Println("stop the goruntine")
	cancel()
	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "got teh stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
