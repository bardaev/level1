package tasks

func Task23() {
	var data []string = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}

	var elem int = 5

	PrintInputOutput(23.1, "Удаление элемента из слайса с нарушением порядка (вариант 1)", removeItem1(data, elem), data)
	PrintInputOutput(23.2, "Удаление элемента из слайса с сохранением порядка (вариант 2)", removeItem2(data, elem), data)
}

// Меняем последний елемент с тем который хотим удалить
// и возвращаем массив с длинной -1
func removeItem1(arr []string, i int) []string {
	var tmp []string = make([]string, len(arr))
	copy(tmp, arr)
	tmp[i] = tmp[len(tmp)-1]
	return tmp[:len(tmp)-1]
}

// Клеим 2 части массива без удаляемого элемента
func removeItem2(arr []string, i int) []string {
	var tmp []string = make([]string, len(arr))
	copy(tmp, arr)
	result := tmp[:i]
	result = append(result, tmp[i+1:]...)
	return result

	// result := arr[:i]
	// result = append(result, arr[i+1:]...)
	// return result
}
