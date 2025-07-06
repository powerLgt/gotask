package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Mutex   sync.Mutex
	Counter int
}

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func main() {
	counter := &Counter{}
	channel := make(chan int)
	for i := 0; i < 10; i++ {
		go func(counter *Counter, channel chan int) {
			for i := 0; i < 1000; i++ {
				counter.Mutex.Lock()
				counter.Counter += 1
				counter.Mutex.Unlock()
			}
			channel <- 1
		}(counter, channel)
	}

	for i := 0; i < 10; i++ {
		<-channel
	}
	fmt.Println("counter: ", counter.Counter)
}
