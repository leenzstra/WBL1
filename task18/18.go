package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Счетчик с мьютексом
type Counter struct {
	V   int
	mut sync.Mutex
}

// блокирующий инкремент
func (c *Counter) Inc() {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.V++
}

// Запуск инкремента incs раз
func runCounter(incs int, c *Counter) {
	for i := 0; i < incs; i++ {
		c.Inc()
	}
}

func main() {
	// Кол-во горутин
	workers := 10

	// Множитель кол-ва инкрементов Inc()
	incsMul := 10

	wg := sync.WaitGroup{}

	counter := Counter{}

	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		// Кол-во вызовов Inc()
		incs := i * incsMul

		go func() {
			runCounter(incs, &counter)
			wg.Done()
		}()
	}

	wg.Wait()

	// Сколько должно получится
	mustResult := (incsMul + incsMul*workers) * workers / 2
	fmt.Printf("must %d, counter %d, %t", mustResult, counter.V, mustResult == counter.V)
}
