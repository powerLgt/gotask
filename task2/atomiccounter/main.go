package main

import (
	"fmt"
	"sync/atomic"
)

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func main() {
	var counter int64
	channel := make(chan int)

	for i := 0; i < 10; i++ {
		go func(counter *int64, channel chan int) {
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(counter, 1)
			}
			channel <- 1
		}(&counter, channel)
	}

	for i := 0; i < 10; i++ {
		<-channel
	}

	fmt.Println("counter: ", atomic.LoadInt64(&counter))
}
