package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Duration(5) * time.Second)
	fmt.Println("开始时间:", time.Now())
	//NewTimer不会堵塞当前的
	for {
		fmt.Println("tewtwtew")
		fmt.Println("执行时间:", <-timer.C)
		//timer.Reset(2 * time.Second)
	}


}
