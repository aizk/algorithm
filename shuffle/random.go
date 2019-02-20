package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 初始化数据
	data = make([]int, numbers)
	for index := range data {
		data[index] = index + 1
	}
	fmt.Printf("%+v", data)

	index := numbers
	for i := 0; i < numbers; i++ {
		// 随即一个下标
		randIndex := rand.Intn(index)
		// 交换
		data[randIndex], data[index - 1] = data[index - 1], data[randIndex]
		// 洗完最后一位
		index--
	}

}

// 牌的数量
const numbers = 52

var data []int
