package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，

	另一个协程从通道中接收这些整数并打印出来。

考察点 ：通道的基本使用、协程间通信。

题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func main() {
	channel := make(chan int)
	channelCached := make(chan int, 100)
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go func(channel chan int, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		for i := 1; i <= 10; i++ {
			channel <- i
			fmt.Println("send channel: ", i)
		}
		close(channel)
	}(channel, &waitGroup)

	go func(channel chan int) {
		for i := range channel {
			fmt.Println("channel receive: ", i)
		}
	}(channel)

	waitGroup.Wait()

	waitGroup.Add(1)

	go func(channel chan int) {
		for i := 1; i <= 100; i++ {
			channel <- i
			fmt.Println("send channel cached: ", i)
		}
		close(channel)
	}(channelCached)

	go func(channel chan int, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		for i := range channel {
			fmt.Println("channel cached receive: ", i)
		}
	}(channelCached, &waitGroup)

	waitGroup.Wait()
	fmt.Println("all done")
}
