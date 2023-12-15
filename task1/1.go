package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Структура Human c публичными полями Name, Age
type Human struct {
	Name string
	Age int
}

// Структура Action со встроенной структурой Human
// Action имеет доступ в полям и методам Human
type Action struct {
	Human
}

// Функция создания нового экземпляра структуры Human
func New(name string, age int) Human {
	return Human{name, age}
}

// Метод структуры Human, выводящий приветствие в консоль
func (h Human) Hello() {
	fmt.Printf("Hello from %s!\n", h.Name)
}


func main() {
	// Создание экзепляра Human
	human := New("Vasyan", 5)

	// Создание экземпляра Action путем встраивания экземпляра Human
	action := Action{human}

	// Вызов метода структуры Human из Action
	action.Hello()
}