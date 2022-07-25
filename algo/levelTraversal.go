package algo

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func levelTraversal(root *Node) {
	if root == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}
