package main

import (
	"fmt"
	"sync"
)

func square(x int) int {
	return x * x
}

func WgSolution(arr [5]int) {
	fmt.Println("WG Solution:")
	wg := sync.WaitGroup{}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		//чтобы горутины не считывали i из общего контекста.
		//Избегаем этой ситуации, передавая arr[i] горутине как аргумент
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(square(n))
		}(arr[i], &wg)
	}
	//нам нужно дождаться выполнения всех горутин, поэтому используем WaitGroup
	wg.Wait()
}

func MUSolution(arr [5]int) {
	fmt.Println("Mu Solution:")
	res := []int{}
	wg := sync.WaitGroup{}

	mu := sync.Mutex{}
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		go func(i int) {
			defer wg.Done()
			//может возникнуть гонка данных, посколько в res будут конкурентно писать несколько горутин.
			// чтобы в единый момент времени в res записывала только одна горутина - используем mutex.
			mu.Lock()
			res = append(res, square(arr[i]))
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(res)

}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	WgSolution(arr)
	MUSolution(arr)
}
