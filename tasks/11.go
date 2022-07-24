package tasks

func Task11() {
	var testData1 []string = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var testData2 []string = []string{"k", "b", "m", "c", "f", "g", "x"}

	var testData3 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var testData4 []int = []int{2, 30, 5, 9, 1, 50, 7, 10, 2}

	PrintInputOutput(
		11.1,
		"Пересечение множеств строк вариант 1 (С помощью структуры map)",
		intersectionMap(testData1, testData2),
		testData1,
		testData2,
	)

	PrintInputOutput(
		11.2,
		"Пересечение множеств чисел вариант 1 (С помощью структуры map)",
		intersectionMap(testData3, testData4),
		testData3,
		testData4,
	)

	PrintInputOutput(
		11.3,
		"Пересечение множеств строк вариант 2 (Перебором)",
		forEach(testData1, testData2),
		testData1,
		testData2,
	)

	PrintInputOutput(
		11.4,
		"Пересечение множеств чисел вариант 2 (Перебором)",
		forEach(testData3, testData4),
		testData3,
		testData4,
	)
}

// Пересечение множеств с помощью мапы
func intersectionMap[T comparable](set1 []T, set2 []T) []T {
	// Сюда запишем уникальные пересечения
	resultMap := make(map[T]T)
	hashMap := make(map[T]T)
	result := make([]T, 0)

	// Копируем из первого множества все значения
	for _, val := range set1 {
		hashMap[val] = val
	}

	// и сравниваем со вторым множеством
	for _, val := range set2 {
		if hashMap[val] == val {
			resultMap[val] = val
		}
	}

	for _, v := range resultMap {
		result = append(result, v)
	}

	return result
}

// Метод обычного перебора
func forEach[T comparable](set1 []T, set2 []T) []T {
	result := make([]T, 0)

	for _, val := range set1 {
		for _, v := range set2 {
			if val == v {
				result = append(result, v)
				// break чтобы были уникальные значения
				break
			}
		}
	}

	return result
}
