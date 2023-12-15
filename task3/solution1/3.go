package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.


func main() {
	// Сумма квадратов
	var sum int64

	// Примитив синхронизации для ожидания завершения группы горутин
	wg := sync.WaitGroup{}

	// Начинаем цикл от 2 до 10 с шагом 2
	for i := 2; i <= 10; i += 2 {
		// Приводим число к  int64
		n := int64(i)
		// Добавляем 1 горутину к ожиданию
		wg.Add(1)

		// Запускаем суммацию в отдельной горутине
		go func(n int64, s *int64) {
			// Добавляем число к сумме с помощью atomic
			atomic.AddInt64(s, n*n)
			// -1 к ожиданию
			wg.Done()
		}(n, &sum)

	}
	// Ожидание выполнения всех горутин
	wg.Wait()

	// Вывод суммы
	fmt.Println(sum)
	fmt.Println(sum == (2 * 2 + 4 * 4 + 6 * 6 + 8 * 8 + 10 * 10))
}