package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Task4() {
	PrintConcurrent(4, "Прерывание воркеров", startTask4)
}

func startTask4() {
	// Количество воркеров
	var workers int = 5
	// Сюда пишем данные
	var ch chan int = make(chan int)
	// Сюда подадим сигнал врайтеру, чтобы остановились воркеры
	var stopWriter chan bool = make(chan bool, 1)

	for i := 0; i < workers; i++ {
		go worker(ch, i)
	}
	go writer(ch, stopWriter)

	stop := make(chan bool)
	signalChan := make(chan os.Signal, 1)
	// Перенаправит CTRL+C в канал
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Println("Process interrupted")
			stopWriter <- true
			time.Sleep(time.Second * 1)
			stop <- true
		}
	}()

	fmt.Println("To stop task press CTRL+C")

	<-stop
}

// В этой горутине мы пишем данные
// Так же из канала мы ждем сигнал стоп
// После получения сигнала стоп мы закроем канал куда писали данные и все воркеры закроются
func writer(ch chan int, stop chan bool) {
	var i int = 0
	for {
		select {
		case <-stop:
			close(ch)
			return
		default:
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker(ch chan int, worker int) {
	for {
		if val, ok := <-ch; ok {
			fmt.Printf("Worker %d: %d\n", worker, val)
		} else {
			break
		}
	}
	fmt.Printf("Worker %d is stopped\n", worker)
}
