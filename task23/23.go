package main

import (
	"fmt"
	"slices"
)

// Удалить i-ый элемент из слайса.

// Удаление сдвигом. Берем первую часть слайса и добавляем в него вторую половину со 2 элемента
func removeByShift(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}

// Заменяем удаляемый элемент на последний (меняется порядок). 
func removeByReplace(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

// Удаление с помощью из стандартного пакета (похож на removeByShift)
func removeWithPackage(s []int, i int) []int {
    return slices.Delete(s, i, i + 1)
}

func main() {
	arr := []int{1,2,3,4}
	arrCopy := make([]int, len(arr))

	copy(arrCopy, arr)
	r := removeByShift(arrCopy, 1)
	fmt.Printf("removeByShift arr: %v copy: %v res: %v\n", arr, arrCopy, r)

	copy(arrCopy, arr)
	r = removeByReplace(arrCopy, 1)
	fmt.Printf("removeByReplace arr: %v copy: %v res: %v\n", arr, arrCopy, r)

	copy(arrCopy, arr)
	r = removeWithPackage(arrCopy, 1)
	fmt.Printf("removeWithPackage arr: %v copy: %v res: %v\n", arr, arrCopy, r)
}