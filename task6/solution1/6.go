package main

import (
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Какие-то вычисления
func someWork() {
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
}

// Пример остановки с помощью канала done
func Run(done <-chan bool) {
	for {
		select {
		// Когда в канал передано значение, выполняется этот кейс
		case <- done:
			return
		default:
			someWork()
		}
	}
}


func main() {
	// Канал для отслеживания остановки горутины
	done := make(chan bool)

	go Run(done)

	time.Sleep(2 * time.Second)

	close(done)
}
