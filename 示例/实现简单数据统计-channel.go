package main

import (
	"fmt"
	"runtime"
	//"sync"
)

var (
	// 数据统计
	counter int
	//wg      sync.WaitGroup
	//mutex   sync.Mutex

	goroutineNum = 6
	forNum       = 2
)

/*
 * 实现简单的调用次数数据统计，
 * 正常的统计结果应该为 goroutineNum*repeatNum = 次 8000
 * 实际运行效果
 */

func main() {
	ch := make(chan int)
	runtime.GOMAXPROCS(6)
	//wg.Add(goroutineNum)
	//go func() {
	for i := 0; i < goroutineNum; i++ {
		go incCounter(i, ch)
	}
	// 接收数据
	for i := 0; i < goroutineNum*forNum; i++ {
		date := <-ch
		counter += date
	}
	fmt.Println("done", counter)
}

// 模拟数据统计
func incCounter(id int, ch chan int) {
	// 看下调用顺序是否与for循环顺序一致
	fmt.Println("ididid", id)
	for count := 0; count < forNum; count++ {
		runtime.Gosched()
		ch <- 1
	}
}
