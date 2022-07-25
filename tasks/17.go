package tasks

import "strconv"

func Task17() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	var target1 int = 2
	var target2 int = 12
	var target3 int = 90

	PrintInputOutput(17.1, "Бинарный поиск", "Позиция: "+strconv.Itoa(binarySearch(arr, target1)), target1)
	PrintInputOutput(17.2, "Бинарный поиск", "Позиция: "+strconv.Itoa(binarySearch(arr, target2)), target2)
	PrintInputOutput(17.2, "Бинарный поиск (несуществующее число)", "Позиция: "+strconv.Itoa(binarySearch(arr, target3)), target3)
}

// Массив должен быть отсортирован
// Сравниваем со значением в середине
// Если меньше то двигаем правую границу влево и наоборот
func binarySearch(slice []int, target int) int {
	var start_idx int = 0
	var end_idx = len(slice) - 1

	for start_idx <= end_idx {
		mid := (start_idx + end_idx) / 2

		if slice[mid] == target {
			return mid
		} else if slice[mid] > target {
			end_idx = mid - 1
		} else {
			start_idx = mid + 1
		}
	}

	return -1
}
