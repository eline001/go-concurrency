package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	// 设置启用的线程数量
	goroutineNum = 16
	// 每个线程执行的的统计次数
	repeatNum = 1000
)

/**
 *  正常的统计结果应该为 goroutineNum*repeatNum = 次 8000
 * 实际运行效果
 */
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		go incCounter(i)
	}
	wg.Wait()
	fmt.Println("done:", counter)
}

// 数据增加
func incCounter(i int) {
	defer wg.Done()
	for count := 0; count < repeatNum; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			// 程序执行到这里，主动退出当前线程，让出cpu
			//fmt.Println("开始让出执行权，我是协程:",i)
			//runtime.Gosched()
			//time.Sleep(time.Second*1)
			//fmt.Println("让出执行权后，我是协程:",i)
			//runtime.Gosched()
			counter++
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}
