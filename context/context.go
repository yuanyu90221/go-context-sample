package context

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Stop the goruntine")
	cancel()
	time.Sleep(5 * time.Second)
}
