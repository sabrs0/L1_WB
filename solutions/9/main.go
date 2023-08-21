package main

import "fmt"

var (
	arr = []int{1, -2, 3, 4, -5, 6, 7, -8, 9, -10, -11, -12, 13, 14, -15}
)

func writeX(arr []int, ch chan int) {
	for _, val := range arr {
		ch <- val
	}
	close(ch)
}

// если будет простой, то виновник - эта горутина
func multX2(ch1, ch2 chan int) {

	for val := range ch1 {
		ch2 <- (val * 2)
	}
	close(ch2)
}

func main() {
	//небуф каналы
	ch1 := make(chan int)
	ch2 := make(chan int)

	go writeX(arr, ch1)
	go multX2(ch1, ch2)
	for ans := range ch2 {
		fmt.Println(ans)
	}

}
