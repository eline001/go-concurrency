package main

import (
	"fmt"
	"time"
)

/**
 * 在死循环中，随机的发送0和1
 */
func main() {
	ch := make(chan int, 1)
	for {
		time.Sleep(time.Second)
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
