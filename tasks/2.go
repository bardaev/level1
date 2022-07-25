package tasks

import "math"

func Task2() {

	var arr []int = []int{2, 4, 6, 8, 10}

	// В данный канал складываем все результаты
	var ch chan int = make(chan int, len(arr))
	defer close(ch)

	begin(arr, ch)

	var task float64 = 2.0

	for i := 0; i < len(arr); i++ {
		task += 0.1
		PrintResult(math.Floor(task*10)/10, "Конкурентный расчет квадратов чисел", <-ch)
	}
}

func begin(arr []int, ch chan int) {

	// Запускаем горутины в количестве
	for i := 0; i < len(arr); i++ {
		go pow(arr[i], ch)
	}

}

func pow(val int, ch chan int) {
	// Отправляем данные в канал
	ch <- int(math.Pow(float64(val), float64(2)))
}
