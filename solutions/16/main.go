package main

import "fmt"

var (
	arr = []int{2, 8, 7, 1, 3, 5, 6, 4}
	//arr = []int{8, 7, 6, 5, 4, 3, 2, 1}
)

func Partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i = i + 1
			A[j], A[i] = A[i], A[j]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
func myQuickSort(A []int, p, r int) []int {
	if p < r {
		q := Partition(A, p, r)
		myQuickSort(A, p, q-1)
		myQuickSort(A, q+1, r)
	}
	return A
}
func main() {
	newArr := myQuickSort(arr, 0, len(arr)-1)
	fmt.Println(newArr)

}
