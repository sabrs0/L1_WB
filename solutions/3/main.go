package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func square(x int) int {
	return x * x
}

func ChanSolution(arr [5]int) {
	fmt.Println("Chan solution:")
	ch := make(chan int, len(arr))
	go func() {
		defer close(ch)
		for i := 0; i < len(arr); i++ {
			ch <- square(arr[i])
		}
	}()
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Println(sum)
}

func MuSolution(arr [5]int) {
	fmt.Println("Mu solution:")
	sum := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		go func(n int) {
			mu.Lock()
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			sum += square(n)
		}(arr[i])
	}
	wg.Wait()
	fmt.Println(sum)
}

func AtomicSolution(arr [5]int) {
	fmt.Println("Atomic solution:")
	var sum int64
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		go func(n int) {
			defer wg.Done()
			atomic.AddInt64(&sum, int64(square(n)))
		}(arr[i])
	}
	wg.Wait()
	fmt.Println(sum)
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	ChanSolution(arr)
	MuSolution(arr)
	AtomicSolution(arr)

}
