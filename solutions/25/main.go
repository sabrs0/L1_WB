package main

import (
	"fmt"
	"time"
)

func sleep1(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C
}
func sleep2(dur time.Duration) {

	<-time.After(dur)
}
func main() {
	dur := time.Second * 5
	fmt.Println("sleep 1:")
	sleep1(dur)
	fmt.Println("sleep 2:")
	sleep2(dur)
}
