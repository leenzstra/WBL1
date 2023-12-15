package main

import (
	"errors"
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(num int64, pos uint, val int) (int64, error) {
	if pos < 0 || pos > 63 {
		return num, errors.New("wrong pos, bounds = [0:63]")
	}
	// создаем маску, устанавливая i бит в 1
	mask := int64(1 << pos)

	if val == 1 {
		// устанавливаем i бит в 1, используя операцию или
		// Пример: 10111010, i = 2
		// 				^
		// маска:  00000100
		// ИЛИ:    10111110
		num = num | mask

	} else {
		// устанавливаем i бит в 0
		//
		// 1. Когда исходный бит = 1, то хватает только XOR (И не изменяет результат)
		// Пример: 10111110, i = 2
		// 				^
		// маска:  00000100
		// XOR:    10111010
		// И:      10111010
		//
		// 2. Когда исходный бит = 0, то нужны XOR и И
		// Пример: 10111010, i = 2
		// 				^
		// маска:  00000100
		// XOR:    10111110
		// И:      10111010 (возвращает к исходному)
		num = (num ^ mask) & num
	}
	return num, nil
}

func main() {
	num := int64(11)
	pos := uint(2)
	val := 1

	r, _ := setBit(num, pos, val)
	
	fmt.Printf("Num: %d, pos: %d, bit: %d, result: %d\n", num, pos, val, r)
	fmt.Printf("Num:\t %064b\n", num)
	fmt.Printf("Result:\t %064b\n", r)
}
