package main

import (
	"fmt"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func main() {
	str := "asd"
	num := 5
	ok := false
	ch := make(chan int)

	values := []interface{}{str, num, ok, ch}

	// Используем type switch. Он берет интерфейс и пытается преобразовать тип
	for _, v := range values {
		switch v.(type) {
		case string:
			fmt.Println(v, "string")
		case bool:
			fmt.Println(v, "bool")
		case int:
			fmt.Println(v, "int")
		case chan int:
			fmt.Println(v, "chan int")
		default:
			fmt.Println(v, "other type")
		}
	}
}