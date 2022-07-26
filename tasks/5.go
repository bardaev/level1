package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task5() {
	PrintConcurrent(5, "Последовательная отправка и чтение канала", starter)
}

func starter() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch chan int = make(chan int)
	// Секунды до завершения
	var deadline int64 = 3
	go producer(ch, &wg, deadline)
	go consumer(ch)

	wg.Wait()
	fmt.Println()
}

// Отправляем данные
// Получаем текущее время и добавляем время через которое завершится продьюсер
// каждые пол секунды отправляем данные
func producer(ch chan int, wg *sync.WaitGroup, deadline int64) {
	end := time.Now().Unix() + deadline

	var i int = 0
	for {
		ch <- i
		i++
		time.Sleep(500 * time.Millisecond)
		if time.Now().Unix() >= end {
			close(ch)
			wg.Done()
			return
		}
	}
}

// Принимаем данные
func consumer(ch chan int) {
	for {
		val, ok := <-ch
		if ok {
			fmt.Print(val)
			fmt.Print(" ")
		} else {
			break
		}
	}
}
