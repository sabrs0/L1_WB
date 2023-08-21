package main

import (
	"sync"
	"time"

	"fmt"
)

var (
	persons = []string{
		"Ivan", "Andrey", "Oleg", "Sergey", "Georgy",
	}
	ages = []int{
		20, 30, 40, 50, 60,
	}
)

type Person struct {
	Name string
	Age  int
}

// запись в обычную map с использованием mutex, для безопасной записи.
func writeUsualMap(p Person, personAges map[string]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	personAges[p.Name] = p.Age
	mu.Unlock()
}

// запись в sync map без использования mutex, есть встроенный у sync map
func writeSyncMap(p Person, personAges *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	personAges.Store(p.Name, p.Age)
}

func UsualMap() {
	personAges := make(map[string]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(persons))
	for i := 0; i < len(persons); i++ {
		go writeUsualMap(Person{
			Name: persons[i],
			Age:  ages[i],
		}, personAges, &wg, &mu)
	}
	wg.Wait()
	fmt.Println("USUAL MAP IS: \n", personAges)
}
func SyncMap() {
	personAges := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(len(persons))
	for i := 0; i < len(persons); i++ {
		go writeSyncMap(Person{
			Name: persons[i],
			Age:  ages[i],
		}, &personAges, &wg)
	}
	wg.Wait()
	fmt.Println("SYNC MAP IS:")
	personAges.Range(func(key, value any) bool {
		fmt.Printf("[%s]%d\n", key, value)
		return true
	})
}

func main() {
	UsualMap()
	time.Sleep(time.Second * 5)
	SyncMap()
}
