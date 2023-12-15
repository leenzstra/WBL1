package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

// Бианрный поиск (v - искомый элементо, arr - отсортированный массив) 
// Возвращает индекс элемента и флаг "найден"
func binarySearch(v int, arr []int) (int, bool) {

	// Задаем начальные границы
	left, right := 0, len(arr)-1

	// Пока валидные границы
	for left <= right {
		// Получеаем индекс и значение центрального элемента между границами
		mid := (left + right) / 2
		cur := arr[mid]

		// Если центральный элемент меньше искомого, сдвигаем границу вправо
		if cur < v {
			left = mid + 1
		// Если центральный элемент больше искомого, сдвигаем границу влево
		} else if cur > v {
			right = mid - 1
		// Иначе они равны, значит элемент найден
		} else {
			return mid, true
		}
	}

	return 0, false

}

func main() {
	arr := []int{1,4,5,8,9,12,}

	// Проверим существование
	valExists := 9 

	// Проверим несуществование
	valNotExists := 6

	idx, ok := binarySearch(valExists, arr)
	fmt.Printf("%d found? %t; idx: %d\n", valExists, ok, idx)

	//  		  V
	//    1 4 5 8 9 12
	// 1  L   ^     R
	// 2        L ^ R -> найден

	idx, ok = binarySearch(valNotExists, arr)
	fmt.Printf("%d found? %t; idx: %d\n", valNotExists, ok, idx)

	//  		             V
	//    1 4 5 8 9 12       6
	// 1  L   ^     R
	// 2        L ^ R
	// 3       L^R
	// 4      R L      -> right < left -> не найден
}