package main

import (
	"fmt"
	"math/rand"
	"time"
)
// 共享map
type MapStr map[string]string

/**
 * 并非情況下，對原生的map操作
 */
func main() {
	// 写数据
	test := MapStr{}
	test.set("test", "test")
	// 设置
	for i := 0; i < 10; i++ {
		go test.set(randString(), randString())
	}
	// 获取
	for i := 0; i < 10; i++ {
		value := test.get("test")
		fmt.Println("value is :", value)
	}
}

// 随机获取字符串
func randString() string {
	//time.Sleep(time.Second * 2)
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	startIndex := rand.Intn(len(char))
	endIndex := rand.Intn(len(char))
	if startIndex > endIndex {
		startIndex, endIndex = endIndex, startIndex
	}
	return char[startIndex:endIndex]
}

func (m *MapStr) set(key string, value string) {
	test := *m
	test[key] = value
}

func (m *MapStr) get(key string) string {
	test := *m
	return test[key]
}
