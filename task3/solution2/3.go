package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

// concurrent safe сумма
type Sum struct {
	value int64
	mut sync.Mutex
}

// Добавить к значению суммы
func (s *Sum) Add(v int64) {
	defer s.mut.Unlock()
	s.mut.Lock()
	s.value += v
}

// Получить значение суммы
func (s *Sum) Value() int64 {
	return s.value
}

func main() {
	// Сумма квадратов
	sum := Sum{}

	// Примитив синхронизации для ожидания завершения группы горутин
	wg := sync.WaitGroup{}

	// Начинаем цикл от 2 до 10 с шагом 2
	for i := 2; i <= 10; i += 2 {

		// Приводим число к типу int64
		n := int64(i)
		// Добавляем +1 дельту к ожиданию
		wg.Add(1)

		// Запускаем суммацию в отдельной горутине
		go func(n int64) {
			// Добавляем число к сумме с помощью atomic
			sum.Add(n*n)
			// -1 к ожиданию
			wg.Done()
		}(n)

	}
	// Ожидание выполнения всех горутин
	wg.Wait()

	// Вывод суммы
	fmt.Println(sum.Value())
	fmt.Println(sum.Value() == (2 * 2 + 4 * 4 + 6 * 6 + 8 * 8 + 10 * 10))
}
