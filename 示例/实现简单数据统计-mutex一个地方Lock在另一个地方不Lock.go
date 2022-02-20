package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
	// 设置启用的线程数量
	goroutineNum = 16
	// 每个线程执行的的统计次数
	repeatNum = 1000
)

/**
 * 如果在一个地方Lock()，在另一个地方不Lock()而是直接修改或访问共享数据，是否可行 ?
 * 这对于sync.Mutex类型来说是允许的，因为mutex不会和goroutine进行关联,
 * 但是他的结果会和我们预期的一样吗？
 */

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(goroutineNum * 2)
	for i := 0; i < goroutineNum; i++ {
		go incCounter()
		go noLockIncCounter()
	}
	wg.Wait()
	fmt.Println(counter)
}

// 数据增加 有lock
func incCounter() {
	defer wg.Done()
	for count := 0; count < repeatNum; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()

		{
			runtime.Gosched()
			counter++
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}

// 不加锁
func noLockIncCounter() {
	defer wg.Done()
	for count := 0; count < repeatNum; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		//mutex.Lock()
		{
			//runtime.Gosched()
			counter++
		}
		//mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}
