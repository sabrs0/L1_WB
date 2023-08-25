package main

import (
	"fmt"
	"strings"
)

var (
	str1 = "abcd"
	str2 = "abCdefAaf"
	str3 = "aabcd"
)

func isUnique(str string) bool {
	lowerStr := strings.ToLower(str)
	runes := []rune(lowerStr)
	uniques := make(map[rune]struct{})
	for _, val := range runes {
		if _, isTaken := uniques[val]; !isTaken {
			uniques[val] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUnique(str1))
	fmt.Println(isUnique(str2))
	fmt.Println(isUnique(str3))
}
