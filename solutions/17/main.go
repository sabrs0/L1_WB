package main

import "fmt"

var (
	arr = []int{2, 8, 7, 1, 3, 5, 6, 4}
)

func myBinSearch(arr []int, elem int) int {

	left := 0
	right := len(arr) - 1
	for left <= right {
		midInd := (left + right) / 2
		midVal := arr[midInd]
		if midVal > elem {
			right = midInd - 1
		} else if midVal < elem {
			left = midInd + 1
		} else {
			return midInd
		}
	}
	return -1
}
func main() {
	ind := myBinSearch(arr, 1)
	if ind == -1 {
		fmt.Println("NOT IN ARRAY")
	} else {
		fmt.Println(arr[ind])
	}
}
