package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

// Свап через присвоение кортежа
func SwapWithTuple(num1, num2 int) {
	was := fmt.Sprintf("%d, %d", num1, num2)

	num1, num2 = num2, num1
	
	fmt.Printf("%s -> %d, %d\n", was, num1, num2)
}

// Свап через мат операции
func SwapWithOps(num1, num2 int) {
	was := fmt.Sprintf("%d, %d", num1, num2)

    num1 = num1-num2
    num2 = num1+num2
    num1 = num2-num1
	
	fmt.Printf("%s -> %d, %d\n", was, num1, num2)
}

func main() {
	n1, n2 := 50, -25
	SwapWithTuple(n1, n2)
	SwapWithOps(n1, n2)
}