package main

import (
	"go-laboratory/algorithm/list"
	"fmt"
)

func main() {
	RecursePrintList(list.TestList)
}

func RecursePrintList(n *list.Node) {
	if n.Next == nil {
		fmt.Print(n.Value)
		return
	}
	RecursePrintList(n.Next)
	fmt.Print(n.Value)
}

// 1. beauty-of-program 各种题目的解法收集
// 2. 基本数据结构实现的轮子
// 3. go-leetcode leetcode