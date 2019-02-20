package main

import (
	. "go-laboratory/algorithm/tree/tree"
	"fmt"
)
// 遍历
func main() {
	fmt.Println(BSTSearch(TestTree, 0))
}


// search
func BSTSearch(n *Node, x int) *Node {
	switch {
	case n == nil:
		return nil
	case n.Key == x:
		return n
	case n.Key > x:
		return BSTSearch(n.Left, x)
	case n.Key < x:
		return BSTSearch(n.Right, x)
	}
	return nil
}