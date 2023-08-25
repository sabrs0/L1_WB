package main

import "fmt"

var (
	set1 = []int{123, 456, 789, -10, 11, -12, 13, 13, 123, -456, 11, 11, 11}
	set2 = []int{987, 654, 321, 789, 111, 13, 11, 11, -10, -10, 13}
)

func intersect(set1, set2 []int) []int {
	elems := make(map[int]int, len(set1))
	for i := 0; i < len(set1); i++ {
		elems[set1[i]]++
	}
	ans := []int{}
	for i := 0; i < len(set2); i++ {
		key := set2[i]
		if _, isOK := elems[key]; isOK {
			if elems[key] > 0 {
				elems[key]--
				ans = append(ans, key)
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(intersect(set1, set2))
}
