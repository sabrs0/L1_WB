package main

import (
	"fmt"
	"sync"
)

var N = 100000

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) MyInc(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(N)
	c := Counter{}
	for i := 0; i < N; i++ {
		go c.MyInc(&wg)

	}
	wg.Wait()
	fmt.Println(c.val)
}
