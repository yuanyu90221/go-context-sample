package channel

import (
	"fmt"
	"time"
)

func Channel() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}

		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goruntine")
	stop <- true
	time.Sleep(5 * time.Second)
}
