package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 共享map
type syncMap struct {
	syncMap sync.Map
}

/**
 * 并非情況下，對原生的map操作
 */
func main() {
	// 写数据
	test := syncMap{}
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

	// 需要获取 map中的所有键值对
	// 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，
	// 返回 true，终止迭代遍历时，返回 false。
	test.syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, "=>", value)
		return true
	})
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

func (m *syncMap) set(key string, value string) {
	m.syncMap.Store(key, value)
	//test.SyncMapStr[key] = value
}

func (m *syncMap) get(key string) string {
	value, _ := m.syncMap.Load(key)
	return value.(string)
}
