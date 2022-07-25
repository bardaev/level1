package tasks

func Task16() {
	var inputData []int = []int{100, 55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79}

	var result []int = make([]int, len(inputData))
	copy(result, inputData)

	PrintInputOutput(16, "Быстрая сортировка", quickSort(result, 0, len(result)-1), inputData)
}

// Функция partiotion сортирует элементы относительно опорной точки
// в левую часть все что меньше в правую все что больше
// и так рекурсивно
func quickSort(slice []int, left int, right int) []int {
	if len(slice) < 1 {
		return slice
	}

	if left < right {
		var p int
		slice, p = partition(slice, left, right)
		slice = quickSort(slice, left, p-1)
		slice = quickSort(slice, p+1, right)
	}

	return slice
}

// Берем опорную точку
// и меняем местами элементы которые меньше опорной с теми которые больше опорной точки
func partition(slice []int, left int, right int) ([]int, int) {
	pivot := slice[right]
	i := left

	for j := left; j < right; j++ {
		if slice[j] < pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}

	slice[i], slice[right] = slice[right], slice[i]
	return slice, i
}
