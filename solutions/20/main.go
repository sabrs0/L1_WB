package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	s1 := strings.Split(str, " ")
	ans := strings.Builder{}
	for i := len(s1) - 1; i >= 1; i-- {
		ans.WriteString(s1[i])
		ans.WriteRune(' ')
	}
	ans.WriteString(s1[0])
	return ans.String()
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
