package tasks

import (
	"sync"
	"sync/atomic"
)

func Task18() {
	PrintResult(18, "Конкурентный счетчик", concurrentWrite())
}

type counter struct {
	count int32
}

// используем atomic, чтобы инкремент выполнялся атомарно
func (c *counter) Add(wg *sync.WaitGroup) {
	atomic.AddInt32(&c.count, 1)
	wg.Done()
}

func concurrentWrite() *counter {
	counter := counter{count: int32(0)}

	n := 100

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			counter.Add(&wg)
		}()
	}

	wg.Wait()

	return &counter
}
