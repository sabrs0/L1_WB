package main

import (
	"math/big"

	"fmt"
)

func Calc(num1, num2 big.Int, operator string) big.Int {
	var ans big.Int
	switch operator {
	case "+":
		ans.Add(&num1, &num2)
	case "-":
		ans.Sub(&num1, &num2)
	case "*":
		ans.Mul(&num1, &num2)
	case "/":
		ans.Div(&num1, &num2)
	default:
		fmt.Println("UNKNOWN OPERATOR")
	}
	return ans
}
func main() {

	num1, _ := new(big.Int).SetString(
		"243214312443214899284509182344321", 10,
	)
	num2, _ := new(big.Int).SetString(
		"1234567890123235543254311923454315", 10,
	)
	ans := Calc(*num1, *num2, "*")
	fmt.Println("MUL: ", ans.Text(10))
	ans = Calc(*num1, *num2, "+")
	fmt.Println("SUM: ", ans.Text(10))
	ans = Calc(*num1, *num2, "/")
	fmt.Println("DIV: ", ans.Text(10))
	ans = Calc(*num1, *num2, "-")
	fmt.Println("SUB: ", ans.Text(10))

}
