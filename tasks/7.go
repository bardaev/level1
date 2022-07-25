package tasks

import "sync"

func Task7() {
	start()
	PrintResult(7, "Конкурентная запись в мапу", dataMap)
}

// Общий ресурс
var dataMap map[int]int = make(map[int]int)

func start() {
	var ch chan bool = make(chan bool)
	defer close(ch)

	var mutex sync.Mutex

	for i := 1; i < 5; i++ {
		go writeToMap(&mutex, ch, i)
	}

	// Ждем окончания всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}
}

func writeToMap(mutex *sync.Mutex, ch chan bool, data int) {
	// Блокируем общий ресурс
	mutex.Lock()

	dataMap[data] = data

	mutex.Unlock()
	ch <- true
}
