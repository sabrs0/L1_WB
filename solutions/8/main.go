package main

import (
	"fmt"
)

func getNumber() (int64, error) {
	var num int64
	fmt.Println("Input int64 number:")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}
func getBitIndex() (int, error) {
	var bitInd int
	fmt.Println("Input bit index[0...63]:")
	_, err := fmt.Scanf("%d", &bitInd)
	if err != nil {
		return 0, err
	}
	if bitInd >= 64 || bitInd < 0 {
		return 0, fmt.Errorf("Incorrect bit index")
	}
	return bitInd, nil
}
func getBitValue() (int, error) {
	var bitVal int
	fmt.Println("Input bit number[0 or 1]:")
	_, err := fmt.Scanf("%d", &bitVal)
	if err != nil {
		return 0, err
	}
	if bitVal != 0 && bitVal != 1 {
		return 0, fmt.Errorf("Incorrect bit value")
	}
	return bitVal, nil
}

// если нужно привести бит к нулю - то конъюнкция(И) с нулем в том месте. какой бит надо обнулить
func dropBit(num int64, ind int) int64 {
	var mask int64
	mask = ^(1 << ind)
	return num & mask
}

// если нужно привести бит к единице - то дизъюнкция(ИЛИ) с единицей в том месте. какой бит надо установить
func setBit(num int64, ind int) int64 {
	var mask int64
	mask = (1 << ind)
	return num | mask
}

func main() {

	num, err := getNumber()
	if err != nil {
		panic(err)
	}
	bitInd, err := getBitIndex()
	if err != nil {
		panic(err)
	}
	bitVal, err := getBitValue()
	if err != nil {
		panic(err)
	}
	var ans int64
	switch bitVal {
	case 0:
		ans = dropBit(num, bitInd)
	case 1:
		ans = setBit(num, bitInd)
	}
	fmt.Println("ANS : ", ans)
}
