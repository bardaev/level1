package tasks

func Task10() {
	var temperature []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	PrintInputOutput(10, "объединение значений в группы", getGroup(temperature), temperature)
}

func getGroup(t []float64) map[int][]float64 {
	var result map[int][]float64 = map[int][]float64{}

	// Проверяем есть ли десятая часть числа
	// Если есть добавляем к слайсу число
	// Если нет создаем слайс с числом
	for _, v := range t {
		if _, ok := result[getStep(v)]; ok {
			result[getStep(v)] = append(result[getStep(v)], v)
		} else {
			result[getStep(v)] = []float64{v}
		}
	}

	return result
}

// Получаем десятую часть числа, то есть ноль на конце
func getStep(num float64) int {
	var step int = 10
	var intNum int = int(num)

	tmp := intNum / step
	return int(tmp) * step
}
