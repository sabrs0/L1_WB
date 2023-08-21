package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func workerStart(ch <-chan int, workerNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("worker %d: handling data - %d\n", workerNumber, val)
		//чтобы после обработки дать возможность работать другим горутинам
		runtime.Gosched()
	}
}

func main() {

	var duration int
	fmt.Println("Input program duraion")
	_, err := fmt.Scanf("%d", &duration)
	if err != nil {
		panic(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(duration))
	//Вызов функции CancelFunc удаляет ссылку родителя на дочерний
	//объект и останавливает все связанные с ним таймеры. Следовательно, после этого
	//сборщик мусора Go может свободно собрать дочерние горутины, у которых больше
	//нет связанных с ними родительских горутин.
	defer cancelFunc()
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go workerStart(ch, 0, &wg)
LOOP:
	for {
		select {
		case <-ctx.Done():
			close(ch)
			break LOOP

		default:
			ch <- rand.Int()
		}
	}
	wg.Wait()

}
