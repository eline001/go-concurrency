package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * 多路复用，接收多个通道的值
 */
func main() {
	// 创建3个通道
	chInt := make(chan int)
	chSting := make(chan string)
	chIntArr := make(chan []int)
	// 新开一个携程
	go func() {
		for i := 0; i < 100000; i++ {
			//statTime := time.Now()

			time.Sleep(time.Second)
			fmt.Println(fmt.Sprintf("第%d次执行", i))
			select {
			case num := <-chInt:
				fmt.Println("收到Int:", num, time.Now())
			case stringResult := <-chSting:
				fmt.Println("收到:StringS,结果为:", stringResult, time.Now())
			case arr := <-chIntArr:
				fmt.Println("收到:arr,结果为:", arr, time.Now())
			}
		}
	}()

	//模拟调用服务（方法，设置不同的时间响应）
	//go func() {
	for {
		time.Sleep(time.Second * 2)
		go sendMsgInt(chInt)
		go sendMsgString(chSting)
		go sendIntArr(chIntArr)
	}
	//}()
}

// 发送整型数据
func sendMsgInt(ch chan int) {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	//time.Sleep(time.Second * 1)
	ch <- rand.Intn(10)
}

// 发送字符串数据
func sendMsgString(ch chan string) {
	//time.Sleep(time.Second * 2)
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	startIndex := rand.Intn(len(char))
	endIndex := rand.Intn(len(char))
	if startIndex > endIndex {
		startIndex, endIndex = endIndex, startIndex
	}
	randString := char[startIndex:endIndex]
	ch <- randString
}

// 发送随机数组
func sendIntArr(ch chan []int) {
	//time.Sleep(time.Second * 3)
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	arrLength := rand.Intn(10)
	var arr []int
	for i := 0; i < arrLength; i++ {
		arr = append(arr, rand.Intn(arrLength))
	}
	ch <- arr
}
