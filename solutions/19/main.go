package main

import "fmt"

func reverseStr(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	s := "главрыба"
	fmt.Println(reverseStr(s))
}
