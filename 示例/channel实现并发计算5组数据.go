package main

import (
	"fmt"
	"math/rand"
)

/**
 * 随机生成5组数据，并且对每组数据求和，最终将5组数据的结果相加
 */
func main() {
	// 随机生成 5组数据,求和
	arrCh := make(chan int)
	go func(arrCh chan int) {
		for i := 0; i < 5; i++ {
			arrNum := rand.Intn(500)
			randArr(arrNum, arrCh)
		}
		// 关闭chanel, 思考一下，如果不关闭channel会不会有问题
		close(arrCh)
	}(arrCh)

	// 所有数据的总和
	var allSum int
	// 通过第四种 循环接受返回结果
	for sumResult := range arrCh {
		fmt.Println("结果为:", sumResult)
		allSum += sumResult
	}
	fmt.Println("所有数据的总和为:", allSum)
}

// 随机生成数组
func randArr(num int, arrCh chan int) {
	var arr []int
	if num > 0 {
		for i := 0; i < num; i++ {
			arr = append(arr, rand.Intn(1000))
		}
	} else {
		arr = append(arr, 0)
	}

	// 计算数据
	arrSum(arr, arrCh)
}

// 计算组数的和
func arrSum(arr []int, arrCh chan int)  {
	var sumResult int
	//fmt.Println("arrarrarrarr", arr)
	for i := 0; i < len(arr); i++ {
		sumResult += arr[i]
	}
	//time.Sleep(time.Second)
	arrCh <- sumResult
}
