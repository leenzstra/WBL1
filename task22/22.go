package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

func main() {
	// Создаем 2 больших числа
	a, b := big.NewInt(1 << 30), big.NewInt(1 << 20)
	// Выделаяем память под результат
	c := new(big.Int)

	fmt.Printf("%d + %d = %d", a, b, c.Add(a, b))
	fmt.Printf("%d - %d = %d", a, b, c.Sub(a, b))
	fmt.Printf("%d * %d = %d", a, b, c.Mul(a, b))
	fmt.Printf("%d / %d = %d", a, b, c.Div(a, b))
}