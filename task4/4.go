package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.


const (
	// Кол-во воркеров по умолчанию
	defaultWorkersCount = 4
		// Кол-во данных
	dataCount = 100000
	// Задержка между отправой данных в миллисекундах
	sendDelay = 0 * time.Millisecond
)

// Функция-воркер, читающая данные из канала
func worker(ctx context.Context, worker int, ch <-chan string, wg *sync.WaitGroup) {
	// При выходе из функции уменьшаем счетчик wg
	defer func() {
		fmt.Println("exit worker", worker)
		wg.Done()
	}()
	
	// В бесконечном цикле проверяем
	// Если пришел сигнал из контекста о завершении, выходим из функции
	// Если пришли данные из канала ch, выводим их
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-ch:
			// Если канал закрыт - скип
			if !ok {
				continue
			}
			fmt.Printf("worker: %d, data: %s\n", worker, data)
		}
	}
}

func main() {
	// Кол-во воркеров из аргументов программы
	workersCount := *flag.Int("n", defaultWorkersCount, "workers count")

	// Контекст с возможностью отмены
	ctx, cancelFunc := context.WithCancel(context.Background())

	// Канал передачи данных в воркеров
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(workersCount)

	// Запускаем воркеров в горутинах
	for i := 0; i < workersCount; i++ {
		go worker(ctx, i, ch, &wg)
	}

	// В отдельной горутине записываем данные в канал ch
	// Если пришел сигнал о завершнии из контекста, то прерываем цикл
	wg.Add(1)
	go func() {
		defer wg.Done()
		loop: 
			for i := 0; i < dataCount; i++ {
			select {
				case <-ctx.Done():
					close(ch)
					break loop
				case ch <- fmt.Sprint(i): 
					time.Sleep(sendDelay)
				
			}
		}
	}()

	// Канал для записи сигналов операционной системы
	termChan := make(chan os.Signal)

	// В канал запишуется сигналы о завершении выполнения программы (interrupt)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Блокирока до тех пор, пока не будет получен сигнал
	<-termChan

	fmt.Println("Shutdown")

	// отменяем контекст, закрываем канал
	cancelFunc()
	// и ожидаем когда все горутины завершатся
	wg.Wait()

	fmt.Print("Exited")
}
