package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

const (
	// Время работы
	defaultRuntime = 2000 * time.Millisecond
	// Задержка между отправкой в канал
	defaultSleep = 0 * time.Millisecond
)

// Функция записи в канал
func writer(ctx context.Context, data chan<- int) {
	// Данные
	i := 0

	// В бесконечном цикле проверяем
	// Если пришел сигнал из контекста о завершении, выходим из функции
	// Если данные успешно записаны в канал, инкремент
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Producer exited")
			return
		case data <- i:
			fmt.Println("Sent", i)
			i++
		}

		time.Sleep(defaultSleep)
	}
}

// Функция чтения из канала
func reader(ctx context.Context, data <-chan int) {
	// В бесконечном цикле проверяем
	// Если пришел сигнал из контекста о завершении, выходим из функции
	// Если пришли данные из канала data, выводим их
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Consumer exited")
			return
		case v, ok := <-data:
			// Если канал закрыт - скип
			if !ok {
				return
			}
			fmt.Println("Received", v, ok)
		}

		time.Sleep(defaultSleep)
	}
}

func main() {
	// Канал передачи данных
	ch := make(chan int)
	// Контекст с отменой
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	// Группа 2 горутин
	wg.Add(2)
	// Чтение из канала
	go func() {
		defer wg.Done()
		reader(ctx, ch)
	}()
	// Запись в канал
	go func() {
		defer wg.Done()
		defer close(ch)
		writer(ctx, ch)
	}()

	// Ожидание N секунд
	time.Sleep(defaultRuntime)

	fmt.Println("Shutdown")
	// отменяем контекст
	cancelFunc()
	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("Exited")

}
