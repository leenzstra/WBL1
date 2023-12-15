package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Какие-то вычисления
func someWork() {
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
}

// Пример остановки с помощью контекста
func Run(ctx context.Context) {
	for {
		select {
		// Когда из контекста истек (сигнал о завершении), выходим из функции
		case <- ctx.Done():
			return
		default:
			someWork()
		}
	}
}

func main() {
	// Контекст с функцией отмены
	// Внутри канал done
	ctx, cancel := context.WithCancel(context.Background())
	go Run(ctx)

	time.Sleep(2*time.Second)

	cancel()
}
