package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task9() {
	PrintConcurrent(9, "Конвеер чисел", startTask9)
}

func startTask9() {
	var wg sync.WaitGroup
	wg.Add(3)

	var dataCh chan int = make(chan int, 10)
	var multDataCh chan int = make(chan int, 10)

	var data []int = []int{1, 2, 3, 4, 5, 6, 7}

	go writeFomArr(dataCh, data, &wg)
	go multNum(dataCh, multDataCh, &wg)
	go readerCh(multDataCh, &wg)

	wg.Wait()
}

// Пишем данные из массива в канал dataCh
func writeFomArr(dataCh chan int, data []int, wg *sync.WaitGroup) {
	for _, val := range data {
		dataCh <- val
		time.Sleep(500 * time.Millisecond)
	}
	close(dataCh)
	wg.Done()
}

// Читаем данные из dataCh и обработанные числа кидаем в канал multDataCh
func multNum(dataCh chan int, multDataCh chan int, wg *sync.WaitGroup) {
	for {
		if val, ok := <-dataCh; ok {
			multDataCh <- val * val
		} else {
			break
		}
	}
	close(multDataCh)
	wg.Done()
}

// Считываем и печатаем данные
func readerCh(multDataCh chan int, wg *sync.WaitGroup) {
	for {
		if val, ok := <-multDataCh; ok {
			fmt.Print(val)
			fmt.Print(" ")
		} else {
			break
		}
	}
	fmt.Println()
	wg.Done()
}
