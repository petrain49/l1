package main

import (
	"fmt"
	"math/big"
)

type nums struct {
	firstValue big.Int
	secondValue big.Int
}

func (n *nums) op(s string) string {
	var res big.Int 

	switch s {
	case "*":
		res.Mul(&n.firstValue, &n.secondValue)
	case "/":
		res.Div(&n.firstValue, &n.secondValue)
	case "+":
		res.Add(&n.firstValue, &n.secondValue)
	case "-":
		res.Sub(&n.firstValue, &n.secondValue)
	default:
		fmt.Println("unknown operation")
	}

	return res.String()
}

func input(firstValue, operation, secondValue string) string {
	nums := nums{}

	nums.firstValue.SetString(firstValue, 10)
	nums.secondValue.SetString(secondValue, 10)

	return nums.op(operation)
}

/*
Разработать программу, которая перемножает, делит,
складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/
func twentyTwo() {
	firstValueString := ""
	secondValueString := ""
	op := ""

	fmt.Scanf("%s %s %s", &firstValueString, &op, &secondValueString)

	fmt.Println(input(firstValueString, op, secondValueString))
}