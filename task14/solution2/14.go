package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func main() {
	str := "asd"
	num := 5
	ok := false
	ch := make(chan int)

	values := []interface{}{str, num, ok, ch}

	// Используем reflect для получения информации о типе
	for _, v := range values {
		fmt.Println(v, reflect.TypeOf(v).String())
	}
}