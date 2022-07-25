package tasks

import (
	"math"
	"math/big"
)

func Task22() {
	var data1 big.Int = *big.NewInt(5000)
	var data2 big.Int = *big.NewInt(2000)

	var data3 big.Int = *new(big.Int)
	data3.SetString("210000000000", 10)
	var data4 big.Int = *big.NewInt(20000000)

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
func isValid(a big.Int, b big.Int) bool {
	var border big.Int = *big.NewInt(int64(math.Pow(2, 20)))
	if a.Cmp(&border) == 1 && b.Cmp(&border) == 1 {
		return true
	}

	return false
}

func mult(a big.Int, b big.Int) *big.Int {
	var result big.Int

	if isValid(a, b) {
		result.Mul(&a, &b)
	}

	return &result
}

func div(a big.Int, b big.Int) *big.Int {
	var result big.Int

	if isValid(a, b) {
		result.Div(&a, &b)
	}

	return &result
}

func plus(a big.Int, b big.Int) *big.Int {
	var result big.Int

	if isValid(a, b) {
		result.Add(&a, &b)
	}

	return &result
}

func minus(a big.Int, b big.Int) *big.Int {
	var result big.Int

	if isValid(a, b) {
		result.Sub(&a, &b)
	}

	return &result
}
