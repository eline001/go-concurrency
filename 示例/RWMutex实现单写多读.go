package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	countRWMutex      int
	countRWMutexGuard sync.RWMutex
	wggR              sync.WaitGroup
	// 并发的数量
	GRN = 10
	// 判断已经设置了多少次
	setNum = 0
)

// 读取数量
func RWGetCount(i int) {
	// 锁定
	//countRWMutexGuard.Lock()
	countRWMutexGuard.RLock()
	time.Sleep(time.Second)
	// 在函数退出时解除锁定
	defer countRWMutexGuard.RUnlock()
	//defer countRWMutexGuard.Unlock()
	fmt.Println(fmt.Sprintf("获取数据[%d], %d, %v", i, countRWMutex, time.Now()))
}

// 设置数量
func RWSetCount(c int) {
	defer wggR.Done()
	countRWMutexGuard.Lock()
	time.Sleep(time.Second*2)
	//time.Sleep(time.Second)
	countRWMutex = c
	setNum++
	fmt.Println("设置数据", c, time.Now())
	// 并发的数量逐渐减一
	countRWMutexGuard.Unlock()
}

func main() {
	// 可以进行并发安全的设置
	fmt.Println("start", time.Now())
	for i := 0; i < GRN; i++ {
		wggR.Add(1)
		go RWSetCount(i)
	}

	// 获取数据(同时申请多个读锁)
	// 可以进行并发安全的获取
	for ; setNum < GRN; {
		time.Sleep(time.Second * 1)
		// 开启两个线程，读取数据
		go RWGetCount(1)
		go RWGetCount(2)
	}
	wggR.Wait()
	fmt.Println("done", countRWMutex, setNum, GRN)
}
