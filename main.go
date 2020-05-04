package main

import (
	channel "github.com/yuanyu90221/go-context-sample/channel"
	context "github.com/yuanyu90221/go-context-sample/context"
	multipleContext "github.com/yuanyu90221/go-context-sample/multiple-context"
)

func main() {
	// var wg sync.WaitGroup

	// wg.Add(2) // 2 job
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("job 1 done.")
	// 	wg.Done()
	// }()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("job 2 done.")
	// 	wg.Done()
	// }()

	// wg.Wait()
	// fmt.Println("All Done.")
	channel.Channel()
	context.Context()
	multipleContext.MultipleContext()
}
