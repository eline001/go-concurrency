package main

import "fmt"

/**
 * 创建一个有缓冲通道，
 */
func main() {
	// 创建一个缓冲区大小为3个的通道
	bufferCh := make(chan int, 3)

	// 在没有另一个线程接受缓冲数据的时候，往里面写数据

	bufferCh <- 1
	fmt.Println("chanLength", len(bufferCh))
	bufferCh <- 2
	fmt.Println("chanLength", len(bufferCh))
	bufferCh <- 3
	fmt.Println("chanLength", len(bufferCh))

	// 1.试想一下，如果我们继续往通道中写数据，会怎么样
	//bufferCh<-4

	// 2. 再试想一下，不断的接受通道的数据
	//for {
	//	fmt.Println("datadatadatadata", <-bufferCh)
	//}
}
