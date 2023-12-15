package main

import (
	"fmt"
	"strings"
)

// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := [5]int{2, 3, 6, 8, 10}

	// Канал для записи квадратов чисел
	results := make(chan int)

	// Запускаем горутину для каждого числа
	for _, n := range arr {
		go func(num int, ch chan<- int) {
			// В канал записываем квадрат числа
			ch <- num * num
		}(n, results)
	}

	// Выделяем память под слайс квадратов
	resultArr := make([]int, 0, len(arr))
	
	// Собираем все результаты в один слайс
	for r := range results {
		resultArr = append(resultArr, r)
		
		// Когда все числа посчитаны, закрываем канал
		if len(resultArr) == len(arr) {
			close(results)
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(resultArr), "[]"))
}
