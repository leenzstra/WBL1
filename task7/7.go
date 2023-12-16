package main

import (
	"fmt"
	"sync"
)

// https://habr.com/ru/articles/338718/
// Реализовать конкурентную запись данных в map.

const (
	// Ключ, по которому происходит запись в map
	key = "key"
	// Кол-во горутин
	workers = 50
)

// concurrent safe map
type ConcMap struct {
	// Мьютекс, разрешающий параллельное чтение и запрещающий параллельную запись
	// При записи используется Lock - Lock и RLock будут ждать разблокировки
	// При чтении используется RLock - Lock ожидает, RLock проходит
	mut sync.RWMutex
	buf map[string]int
}

// Получение значения из мапы по ключу
func (m *ConcMap) Get(key string) (int, bool) {
	// Блокирока для чтения
	// При выходи из функции, разблокировка
	defer m.mut.RUnlock()
	m.mut.RLock()
	data, ok := m.buf[key]
	return data, ok
}

// Установка значения
func (m *ConcMap) Set(key string, value int) {
	// Блокирока для записи
	// При выходи из функции, разблокировка
	defer m.mut.Unlock()
	m.mut.Lock()
	m.buf[key] = value
}

// Добавление к значению по ключу
func (m *ConcMap) Add(key string, value int) {
	// Блокирока для записи
	// При выходи из функции, разблокировка\
	defer m.mut.Unlock()
	m.mut.Lock()
	m.buf[key] += value
}

func main() {
	wg := sync.WaitGroup{}

	m := ConcMap{
		buf: make(map[string]int),
	}

	wg.Add(workers)

	// Каждая горутина прибавит свой номер 
	// в мапу к значению по ключу "key"
	// Пример: workers
	// Результат: 1 + 2 + 3 + 4 + 5 = 15
	for i := 1; i <= workers; i++ {
		go func(n int){
			m.Add(key, n)
			wg.Done()
		}(i)		
	}

	// Ожидание выполнения всех горутин
	wg.Wait()

	// Итоговое значение
	v, _ := m.Get(key)
	// Сумма алг. прогресси от 1 до workers
	sumValue := (1 + workers) * workers / 2

	// Проверка суммы
	fmt.Printf("%d == %d ? %t", v, sumValue, sumValue == v)
}

