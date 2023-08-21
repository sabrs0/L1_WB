package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
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
	var workersN int
	fmt.Println("Input workers num:")
	_, err := fmt.Scanf("%d", &workersN)
	if err != nil {
		panic(err)
	}
	//создаем канал сигналов, в который будет приходить сигнал, сгенерированный ctrl+C
	sigs := make(chan os.Signal, 1)
	//Описываем сигналы, которые нас интересуют
	signal.Notify(sigs, os.Interrupt)
	wg := sync.WaitGroup{}
	wg.Add(workersN)
	ch := make(chan int, workersN)
	//запускаем воркеров
	for i := 0; i < workersN; i++ {
		go workerStart(ch, i, &wg)
	}
	// создаем метку, чтобы была возможность прервать for внутри select
LOOP:
	for {
		select {
		//обработка сигнала
		case <-sigs:
			/*
				В отличие от переменных, потерянные go-подпрограммы не собираются сборщиком мусора автоматически,
				поэтому важно гарантировать, что go-подпрограммы должны прекратиться сами,
				когда они больше не нужны
			*/
			fmt.Println("Handling ctrl + C")
			close(ch)
			break LOOP
		//отправка данных воркерам
		default:
			ch <- rand.Int()
		}

	}
	wg.Wait()
}
