package main

import (
	"fmt"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {

	// Функция записи последовательности чисел в канал
	// Возвращает канал для чтения
	numPipe := func(nums ...int) <-chan int {
		// Создаем канал
		out := make(chan int)
		// Наполнение канала в новой горутине
		go func() {
			// Записываем числа в канал
			for _, n := range nums {
				out <- n
			}
			// Закрываем канал
			close(out)
		}()
		return out
	}

	// Функция возведения чисел из канала в квадрат
	// Принимает канал чисел
	// Возвращает канал квадратов для чтения
	sqrPipe := func(in <-chan int) <-chan int {
		// Создаем канал квадратов
		out := make(chan int)
		// Наполнение канала в новой горутине
		go func() {
			// Записываем квадраты в канал
			for n := range in {
				out <- n * n
			}
			// Закрываем канал
			close(out)
		}()
		return out
	}

	// Функция вывода чисел 
	// Принимает канал чисел
	stdoutPipe := func(in <-chan int) {
		for n := range in {
			fmt.Println(n)
		}
	}

	// Все горутины работают паралелльно - main, sqr, num

	stdoutPipe(sqrPipe(numPipe(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))
}
