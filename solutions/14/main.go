package main

import (
	"fmt"
	"reflect"
)

func switchCheck(intV, stringV, boolV, chanV interface{}) {
	switch intV.(type) {
	case int:
		fmt.Println("TYPE OF ", intV, "IS INT")
	default:
		fmt.Println("TYPE OF ", intV, "IS NOT INT")
	}

	switch stringV.(type) {
	case string:
		fmt.Println("TYPE OF ", stringV, "IS STRING")
	default:
		fmt.Println("TYPE OF ", stringV, "IS NOT STRING")
	}

	switch boolV.(type) {
	case bool:
		fmt.Println("TYPE OF ", boolV, "IS BOOL")
	default:
		fmt.Println("TYPE OF ", boolV, "IS NOT BOOL")
	}

	switch chanV.(type) {
	case chan int:
		fmt.Println("TYPE OF ", chanV, "IS CHAN INT")
	default:
		fmt.Println("TYPE OF ", chanV, "IS NOT CHAN INT")
	}

}
func reflectCheck(intV, stringV, boolV, chanV interface{}) {
	t1 := reflect.TypeOf(intV)
	t2 := reflect.TypeOf(stringV)
	t3 := reflect.TypeOf(boolV)
	t4 := reflect.TypeOf(chanV)
	fmt.Println("TYPE OF ", intV, "IS", t1)
	fmt.Println("TYPE OF ", stringV, "IS", t2)
	fmt.Println("TYPE OF ", boolV, "IS", t3)
	fmt.Println("TYPE OF ", chanV, "IS", t4)
}
func main() {
	var intV, stringV, boolV, chanV interface{}
	intV = 1
	stringV = "Hello"
	boolV = true
	chanV = make(chan int)
	switchCheck(intV, stringV, boolV, chanV)
	fmt.Println("\n\n\n")
	reflectCheck(intV, stringV, boolV, chanV)
}
