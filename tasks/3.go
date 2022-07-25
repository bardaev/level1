package tasks

func Task3() {
	var arr []int = []int{2, 4, 6, 8, 10}
	// В этот канал закидываем квадрат числа
	var ch chan int = make(chan int, len(arr))
	defer close(ch)

	var result int = 0

	// Запускаем горутину которая кидает в канал квадрат числа
	for i := 0; i < len(arr); i++ {
		go square(arr[i], ch)
	}

	// Складываем квадраты чисел
	for i := 0; i < len(arr); i++ {
		result += <-ch
	}

	PrintInputOutput(3, "Сумма квадратов чисел", result, arr)
}

func square(n int, ch chan int) {
	ch <- n * n
}
