package tasks

import (
	"strconv"
	"time"
)

func Task25() {
	var dur1 int64 = Second * 1
	var dur2 int64 = Second * 2
	var dur3 int64 = Second * 3

	start1, end1 := sleep(dur1)
	start2, end2 := sleep(dur2)
	start3, end3 := sleep(dur3)

	PrintInputOutput(25.1, "", "Start: "+strconv.FormatInt(start1, 10)+", end: "+strconv.FormatInt(end1, 10), strconv.FormatInt(dur1, 10)+" seconds")
	PrintInputOutput(25.2, "", "Start: "+strconv.FormatInt(start2, 10)+", end: "+strconv.FormatInt(end2, 10), strconv.FormatInt(dur2, 10)+" seconds")
	PrintInputOutput(25.3, "", "Start: "+strconv.FormatInt(start3, 10)+", end: "+strconv.FormatInt(end3, 10), strconv.FormatInt(dur3, 10)+" seconds")
}

// Для реализации time.Sleep берем текущее время и добавляем количество секунд, которое будет означать конец функции sleep
// в бесконечном цикле проверяем текущее время и сравниваем с переменной end
// когда время будет равно с конечным завершаем функцию
func sleep(duration int64) (int64, int64) {
	var start int64 = time.Now().Unix()
	var end int64 = time.Now().Unix() + duration

	for {
		if time.Now().Unix() == end {
			return start, end
		}
	}
}

const (
	Second int64 = 1
	Minute int64 = Second * 60
	Hour   int64 = Minute * 60
)
