package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

// В задании не указано сколько может быть пробелом, поэтому пусть всегда будет 1

// // inplace функция
// func ReverseInplace(v []string) {
// 	// Меняем попарно значения с начала и конца
// 	for l, r := 0, len(v)-1; l < r; l, r = l+1, r-1 {
// 		v[l], v[r] = v[r], v[l]
// 	}
// }

// Реверсирование массива
func Reverse(v []string) []string {
	// Выделяем память под перевернутую версию
	revValues := make([]string, len(v))

	// Перемещаем значения с конца в начало
	for i := 0; i < len(v); i++ {
		revValues[i] = v[len(v)-i-1]
	}

	return revValues
}

// Разбиение строки v разделителем div
func Split(v, div string) []string {
	return strings.Split(v, div)
}

// Реверсирование слов в строке v,
// div - разделитель слов
func ReverseWords(v, div string) string {
	return strings.Join(Reverse(Split(v, div)), div)
}

func main() {
	s := "snow dog sun"
	revWords := ReverseWords(s, " ")

	fmt.Printf("%s > %s", s, revWords)

}
