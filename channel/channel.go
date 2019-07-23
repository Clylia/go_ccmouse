package main

import (
	"fmt"
	"time"
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

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worked %d received %c\n", id, <-c)
	}
}

// channel demo
func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}

func main() {
	//channelDemo()
	bufferedChannel()
}
