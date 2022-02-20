package main

import (
	"fmt"
	"net/http"
	"sync"
)

/**
* 通过 sync.WaitGroup 等待，任务执行完成，然后再退出主进程
 * 假如没有 sync.WaitGroup; 我们看下，结果会怎么样？
*/
func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		// 开启一个并发
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}
