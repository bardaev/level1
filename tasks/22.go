package tasks

import "math"

func Task22() {
	var data1 float64 = 50000
	var data2 float64 = 50000

	var data3 float64 = 1200000
	var data4 float64 = 1500000

	PrintInputOutput(22.1, "Умножение чисел меньше 2^20", mult(data1, data2), data1, data2)
	PrintInputOutput(22.2, "Умножение чисел больше 2^20", mult(data3, data4), data3, data4)
	PrintInputOutput(22.3, "Деление чисел меньше 2^20", div(data1, data2), data1, data2)
	PrintInputOutput(22.4, "Деление чисел больше 2^20", div(data3, data4), data3, data4)
	PrintInputOutput(22.5, "Сложение чисел меньше 2^20", plus(data1, data2), data1, data2)
	PrintInputOutput(22.6, "Сложение чисел больше 2^20", plus(data3, data4), data3, data4)
	PrintInputOutput(22.5, "Вычитание чисел меньше 2^20", minus(data1, data2), data1, data2)
	PrintInputOutput(22.6, "Вычитание чисел больше 2^20", minus(data3, data4), data3, data4)
}

// Везде проверяем что оба числа больше 2^20
func isValid(a float64, b float64) bool {
	if a > math.Pow(2, 20) && b > math.Pow(2, 20) {
		return true
	}

	return false
}

func mult(a float64, b float64) float64 {
	var result float64

	if isValid(a, b) {
		result = a * b
	}

	return result
}

func div(a float64, b float64) float64 {
	var result float64

	if isValid(a, b) {
		result = a / b
	}

	return result
}

func plus(a float64, b float64) float64 {
	var result float64

	if isValid(a, b) {
		result = a + b
	}

	return result
}

func minus(a float64, b float64) float64 {
	var result float64

	if isValid(a, b) {
		result = a - b
	}

	return result
}
