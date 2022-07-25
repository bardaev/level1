package tasks

import (
	"context"
	"sync"
	"time"
)

func Task6() {
	// Variant 1
	n := 100
	var wg sync.WaitGroup
	wg.Add(n)

	var ch chan int = make(chan int, n)

	for i := 0; i < n; i++ {
		go stopVariant1(&wg, ch)
	}

	for i := 0; i < n; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

	PrintResult(6.1, "Остановка горутины вариант 1", nil)

	// Variant 2
	channel := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go stopVariant2(ctx, channel)
	go durationStop(cancel)

	<-channel
	PrintResult(6.2, "Остановка горутины вариант 2", nil)
}

// Запускаем 100 горутин
// Посылаем туда данные и закрываем канал
// При закрытии канала горутина завершается
func stopVariant1(wg *sync.WaitGroup, ch chan int) {
	for {
		_, ok := <-ch
		if !ok {
			wg.Done()
			return
		}
	}
}

// Используеи пакет context
// когда вызывается cancel() то возвращается канал в ctx.Done(), который сигнализирует об отмене context
// после этого мы прерываем функцию
func stopVariant2(ctx context.Context, ch chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			ch <- struct{}{}
			return
		default:

		}
		time.Sleep(500 * time.Millisecond)
	}
}

func durationStop(cancel context.CancelFunc) {
	time.Sleep(3 * time.Second)
	cancel()
}
