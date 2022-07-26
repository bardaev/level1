package tasks

import (
	"strings"
)

func Task15() {
	var length int = 100

	PrintResult(15, "Задача копирования слайса", someFunc(length))
}

// Проблема данной функции в том, что выделяется память под большую строку
// из нее берется часть слайса. Этот слайс меньшей длины чем исходная строка
// но эти данные указывают на исходную строку. Эта ссылка хрунится в памяти
// и GC не отчистит память от этой строки пока на нее есть ссылка

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

var justString string

// Чтобы избежать этого. Необходимо скопировать часть строки при помощи функции copy
// Тогда больше ничего не будет указывать на большую строку и память отчистится от нее
// и у нас будет в памяти скопированная строка меньшей длины, которая не ссылается на большую строку
func someFunc(length int) string {
	v := createHugeString(1 << 10)
	newStr := make([]rune, length)
	copy(newStr, []rune(v[:length]))
	justString := newStr
	return string(justString)
}

func createHugeString(length int) string {
	return strings.Repeat("a", length)
}
