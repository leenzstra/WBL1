package main

import (
	"fmt"

	"github.com/leenzstra/WBL1/task24/point"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.


func main() {
	p1 := point.NewPoint(1, 2)
	p2 := point.NewPoint(2, -5)

	d := p1.Distance(p2)

	fmt.Printf("p1: %s, p2: %s, dist: %f", p1, p2, d)
}
