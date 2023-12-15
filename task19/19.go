package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

func Reverse(v string) string {
	// Дастаем руны из строки
	rs := []rune(v)

	// Меняем попарно значения с начала и конца
	for l, r := 0, len(rs)-1; l < r; l, r = l+1, r-1 {
        rs[l], rs[r] = rs[r], rs[l]
    }

	// Преобразуем руны в строку
	return string(rs)
}

func main() {
	s := "a狐 фыв 狐b"
	reversed := Reverse(s)

	fmt.Printf("%s > %s", s, reversed)
}