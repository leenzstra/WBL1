package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.


// Функция создания множества из массива элементов.
// Используется map для определения уникальности элементов.
// Каждый элемент массива представляется ключом в map.
// Массив ключей является множеством 
func newSet(words []string) []string {
	// Выделяем память под ключи
	keys := make([]string, 0)
	
	// Выделяем память под мапу
	setMap := make(map[string]bool)

	// Добавляем все элементы в мапу
	for _, s := range words {
		setMap[s] = true
	}

	// Получем множество ключей
	for k := range setMap {
		keys = append(keys, k)
	}

	return keys
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	keys := newSet(seq)
	fmt.Println(keys)
}