package main

import (
	"fmt"
	//"runtime"

	//"runtime"
	"time"

	//"runtime"
	//"sync"
	//"time"

	//"runtime"

)

func say(s string) {
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Second*2)
		//runtime.Gosched()
		time.Sleep(time.Millisecond*500)
		fmt.Println(s)
	}
	//wg.Done()
}

//var wg sync.WaitGroup

func main() {
	startTime := time.Now()
	//wg.Add(2)
	 say("world")
	 say("hello")
	 say("你好")
	 say("世界")
	//wg.Wait()
	endTime := time.Now()
	fmt.Println(fmt.Sprintf("执行结束\n，开始执行时间:%v, 执行结束时间：%v,总的花了%v", startTime, endTime, endTime.Sub(startTime)))
}
