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
	for n := range c {
		fmt.Printf("Worked %d received %d\n", id, n)
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
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	channelDemo()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
