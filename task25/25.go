package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func uneffectiveSleep(d time.Duration) {
	// Текущее время
	now := time.Now()
	// Время в будущем, которое ожидаем
	final := now.Add(d)

	// Проверяем, что будущее еще не настало
	if final.Before(now) {
		return
	}

	// Ждем будущее в цикле, каждый шаг проверяя оставееся время
	for final.After(time.Now()) {
		
	} 
}

func main() {
	fmt.Println("start")
	uneffectiveSleep(5 * time.Second)
	fmt.Println("slept")
}