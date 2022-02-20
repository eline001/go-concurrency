package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量
var (
	m  sync.Mutex
	v1 int
)

// 修改共享变量
// 在Lock()和Unlock()之间的代码部分是临界区
func change(i int) {
	m.Lock()
	//fmt.Println("我是第:",i,"个")
	// 主动让出执行权
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

// 访问共享变量
// 在Lock()和Unlock()之间的代码部分是是临界区
func read() int {
	m.Lock()
	// 可以试图加个时间间隔，让结果更显而易见
	//time.Sleep(time.Millisecond*200)
	a := v1
	m.Unlock()
	return a
}

/**
 * 观察多次执行结果，从前面0到9个进程的执行结果，和第 11个到19个的变化
 * 当写完数之后，读取数据的方法，并不一定能获取锁成功，所以打印的数据，可能不是顺序的
 */
func main() {
	var numGR = 21
	var wg sync.WaitGroup
	// 第一次打印
	fmt.Printf("%d", read())
	// 循环g创建numGR个oroutine
	// 每个goroutine都执行change()、read()
	// 每个change()和read()都会持有锁
	count :=0
	for i := 0; i < numGR; i++ {
		count++
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("-> 【%d】", read())
		}(i)
	}
	wg.Wait()
}