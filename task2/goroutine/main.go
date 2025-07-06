package main

import (
	"fmt"
	"sync"
	"time"
)

func printOdd(start, end int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := start; i <= end; i++ {
		if i%2 != 0 {
			fmt.Println("find Odd: ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func printEven(start, end int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println("find Even: ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go printOdd(1, 10, &waitGroup)
	go printEven(2, 10, &waitGroup)

	waitGroup.Wait()
}
