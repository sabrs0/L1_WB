package main

import "fmt"

func DeleteFromSlice(arr []int, ind int) ([]int, error) {
	if ind > len(arr)-1 || ind < 0 {
		return nil, fmt.Errorf("Incorrect Index")
	}
	return append(arr[:ind], arr[ind+1:]...), nil
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	res, _ := DeleteFromSlice(arr1, 1)
	fmt.Println(res)

}
