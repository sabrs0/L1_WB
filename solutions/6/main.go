package main

import (
	"context"
	"time"

	"fmt"
)

func ContextCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context Cancelling Gorutine")
			return
		default:
			fmt.Println("ContextCancel Gorutine working")
		}
	}
}

func ChannelCancel(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Channel cancelling gorutine")
			return
		default:
			fmt.Println("ChannelCancel Gorutine working")
		}
	}
}
func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	go ContextCancel(ctx)

	time.Sleep(time.Second * 5)
	ch := make(chan struct{})
	go ChannelCancel(ch)
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
}
