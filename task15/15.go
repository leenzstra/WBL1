package main

import (
	"fmt"
	"strings"
	"unsafe"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// Возможные ошибки:
// 1. 	Создается очень большая строка, а затем берется только ее небольшой срез
// 		при этом срез ссылается на всю область памяти строки (storage)
// 		Из-за этого GC не может ее удалить
// 		

var justString string

func createHugeString(s int) string {
	return strings.Repeat("a", s)
}

func someFunc() {
	// Создается большая строка
	v := createHugeString(1 << 10)

	// ERR: Неэффективно
	justString = v[:100]

	// FIX: Создаем копию в новой области памяти
	justStringNew := strings.Clone(v[:100])

	// Проверка указателей на байты в области памяти 
	fmt.Println(unsafe.StringData(v))				// 0xc000074000
	fmt.Println(unsafe.StringData(justString))		// 0xc000074000
	fmt.Println(unsafe.StringData(justStringNew))	// 0xc000058070

	justString = justStringNew
}

func main() {
	someFunc()
}
