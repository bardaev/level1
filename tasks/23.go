package tasks

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

// С помощью функции append соединяем 2 слайса
// Используется оператор среза n:m
// В первом аргументе передаем слайс от 0 до 4 элемента включительно
// Во втором от 6 и до конца слайса
func Task23() {
	result := append(data[:elem], data[elem+1:]...)

	PrintBeforeAfterResult(23, "Удаление элемента из слайса", data, result)
}
