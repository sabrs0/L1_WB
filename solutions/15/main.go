package main

import (
	"fmt"
	"strings"
)

/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // В нашем примере с символом ♠ вывелось бы 33 символа ♠ и 1 байт из этого символа
}
func main() {
	someFunc()
}*/

func createHugeString(len int) string {
	builder := strings.Builder{}
	for i := 0; i < len; i++ {
		builder.WriteRune('♠') // данный символ весит 3 байта
	}
	return builder.String()
}

// мы получим первые 100 байт сроки, а не 100 символов
// если используется кодировка, в которой символ весит более 1 байта и
// 100 нацело не делится на это число, по последний символ будет утерян, так как
// запишется только его часть.
// если требуется именно 100 символов первых, то:
func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	runeV := []rune(v)
	justString = string(runeV[:100])
	fmt.Println(justString)
}
func main() {
	someFunc()
}
