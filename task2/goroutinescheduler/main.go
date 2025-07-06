package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doSomething() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func calcTime(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	start := time.Now()
	doSomething()
	fmt.Printf("执行耗时: %vms \n", time.Since(start).Milliseconds())
}

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go calcTime(&waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("all done")
}
