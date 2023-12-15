package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

// Проверка уникальности с помощью сортироки
func checkUniqueSort(s string) bool {
	// Создаем массив рун из строки
	runes := []rune(strings.ToLower(s))
	// Сортируем их
	slices.Sort(runes)

	// Попарно сравниваем руны
	// Первая и вторая, вторая и третья, ...
	// Если они равны, то символы неуникальны
	for i := 0; i < len(s) - 1; i++ {
		if runes[i] == runes[i+1] {
			return false
		}
	}

	return true
}

// Проверка уникальности с помощью map (множества)
func checkUniqueMap(s string) bool {
	// Создаем множество, используя map
	runes := make(map[rune]bool)

	// Для каждого символа проверяем
	// Если символ уже есть в множесте, то символы неуникальный
	// Иначе добавляем символ в множество
	for _, r := range strings.ToLower(s) {
		if runes[r] {
			return false
		}
		runes[r] = true
	}

	return true
}

func main() {
	// Примеры для проверки
	examples := []string{"abcd", "", "abCdefAaf", "aabcd", "abcdB"}

	// Запуск двух способов
	// Для каждого определяется время выполнения

	for _, e := range examples {
		start := time.Now()
		r := checkUniqueSort(e)
		fmt.Printf("checkUniqueSort\t example: %s\t result: %t\t time: %s\n", e, r, time.Since(start))
	}

	for _, e := range examples {
		start := time.Now()
		r := checkUniqueMap(e)
		fmt.Printf("checkUniqueMap\t example: %s\t result: %t\t time: %s\n", e, r, time.Since(start))
	}
}