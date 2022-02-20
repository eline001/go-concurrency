package main

import (
	"fmt"
	"time"
)

// 结束标志
/**
 * 生产者
 * 生产10条数据，只有上一条数据被消费之后，下一条数据才可以放入队列中
 * 生产的速度为 1秒每次
 * 0 数字字符表示生产结束
 */

// 结束标志
var endSignMun = 0

func producer(ch chan int, messageNum int, quitCh chan int) {
	for i := 1; i <= messageNum; i++ {
		time.Sleep(time.Second * 1)
		// 将数据通过channel投送给printer
		fmt.Println("P: 我生产了", i)
		ch <- i
	}
	// 0 数字字符表示生产结束
	fmt.Println("P: 我这边生产队已经结束")
	quitCh <- endSignMun
}

/**
 * 消费者
 * 消费的速度为，2秒每次
 */
func consumers(c chan int, quitCh chan int) {
	// 开始无限循环等待数据
	var data int
	for {
		// 从channel中获取一个数据
		time.Sleep(time.Second * 5)
		data := <-c
		// 将0视为数据结束
		if data == endSignMun {
			break
		}
		// 打印数据
		fmt.Println("C:我消费了", data)
	}
	// 通知main已经结束循环(我搞定了!)
	fmt.Println(fmt.Sprintf("C：我接收到了，%d, 累死我了，终于干完活了，该去约会了。。", data))
	quitCh <- endSignMun
}

// 消费者
func main() {
	// 创建一个channel
	c := make(chan int)
	quitCh := make(chan int)
	// 并发执行printer, 传入channel
	go consumers(c, quitCh)
	go producer(c, 10, quitCh)
	//  等待中。。。。。 老板说，你们都干完了，就需要向我汇报一下,收到0，表示两个员工都干完了
	endNum := <-quitCh
	if endSignMun == endNum {
		fmt.Println("B: 你们完成得真棒，每人奖励1个亿")
	} else {
		fmt.Println(fmt.Sprintf("B: 竟敢拿%d来糊弄我，都给我回来，决战到天亮", endNum))
	}
}
