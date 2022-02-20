package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * 模拟任务处理，限定处理的任务不能超过3秒钟
 * 如果超过3秒，则认为程序超时
 */
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	// 新开一个携程
	go func() {
		for {
			statTime := time.Now()
			select {
			case num := <-ch:
				fmt.Println("num=", num)
				//如果超过3秒，则认为程序超时
				// 3秒后执行该操作，条件成立
			case <-time.After(3 * time.Second+time.Millisecond*200):
				endTime := time.Now()
				fmt.Println("quit : time", statTime, "end", endTime)
				quit <- true
			}
		}
	}()

	//模拟调用服务（方法，设置不同的时间响应）
	go func() {
		for {
			// 随机0~4秒
			rand.Seed(time.Now().UnixNano())
			timeSecond := rand.Intn(5)
			fmt.Println("randTime,", timeSecond)
			switch timeSecond {
			case 1:
				time.Sleep(time.Second * 1)
			case 2:
				time.Sleep(time.Second * 2)
			case 3:
				time.Sleep(time.Second * 3)
			case 4:
				time.Sleep(time.Second * 4)
			}
			ch <- timeSecond
		}
	}()

	// 收到退出的信号
	<-quit
	fmt.Println("程序结束")
}
