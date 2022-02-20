package main

import (
	"fmt"
	"net/http"
	"time"
)


/**
 * 假设某个服务，需要请求这些后端服务地址，才能完成某项功能
 * 串行执行
 */
func main() {
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
	startTime := time.Now()
	// 遍历这些地址
	for _, url := range urls {
		http.Get(url)
		// 访问完成后, 打印地址和可能发生的错误
		//fmt.Println(url, err)
	}
	endTime := time.Now()
	// 等待所有的任务完成
	fmt.Println(fmt.Sprintf("串行调用over\n,总的%d个地址\n，开始时间是: %v\n, 结束时间是: %v\n, 总耗时为:%v",len(urls), startTime, endTime, endTime.Sub(startTime)))
}
