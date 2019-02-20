package tree

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var TestTree = &Node{ Key: 1 }

func init() {
	TestTree.Left = &Node{ Key: 0 }
	TestTree.Right = &Node{ Key: 2 }
}

// 中序遍历（也是 BSF 遍历）
// 先打印左子树的键，再打印根结点的键，再打印右子树的键
func Range(node *Node) {
	if node == nil {
		return
	}
	Range(node.Left)
	fmt.Println(node.Key)
	Range(node.Right)
}

// 先序

// 后序