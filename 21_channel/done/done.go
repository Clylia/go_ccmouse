package main

import (
	"fmt"
	"sync"
)

/**
|-------------------------------------------------------------------------------------------------
|
|	channel 并发管道
|
|	c = make(chan int)	// 创建一个 channel
|	c <-	// 接收数据
|	<- c	// 发送数据
|
|-------------------------------------------------------------------------------------------------
*/

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worked %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

// channel demo
func channelDemo() {
	var wg = sync.WaitGroup{}

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	fmt.Println("Channel as first-class citizen")
	channelDemo()
}
