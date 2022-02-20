package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

/**
 * 假设某个服务，需要请求这些后端服务地址，才能完成某项功能
 * 并行执行
 */
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 声明一个等待组
	//var wg sync.WaitGroup
	// 准备一系列的网站地址
	var urls = []string{
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
		"https://www.qiniu.com/",
	}
	// 遍历这些地址
	startTime := time.Now()
	for k, url := range urls {
		//wg.Add(1)
		// 开启一个并发
		go func(url string, k int) {
			//defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("sfsefw", k)
			http.Get(url)
			//fmt.Println(url, err)
		} (url, k)
	}
	// 等待所有的任务完成
	//wg.Wait()
	endTime := time.Now()
	fmt.Println(fmt.Sprintf("串行调用over\n,总的%d个地址\n，开始时间是: %v\n, 结束时间是: %v\n, 总耗时为:%v",len(urls), startTime, endTime, endTime.Sub(startTime)))
}
