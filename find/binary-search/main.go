package main

import "fmt"

func main() {
	// 首先要有序
	fmt.Println(bs([]int{1, 2, 3}, 1))
}

// 二分查找
func bs(data []int, value int) (index int, exist bool) {
	length := len(data)
	if length == 0 {
		return
	}

	low := 0
	height := length - 1
	for low <= height {
		mid := (low + height) / 2
		v := data[mid]
		switch {
		case v == value:
			return mid,true
		case v < value:
			// 注意 + 1 排除掉中间这个数字
			low = mid + 1
		case v > value:
			// 注意 - 1 排除掉中间这个数字
			height = mid - 1
		}
	}
	return
}