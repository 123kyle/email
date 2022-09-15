package code09

import (
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isFull(head *Node) bool {
	if head == nil {
		return true
	}
	height, count := process(head)

	return int(math.Pow(2.0, float64(height))) == count+1
}

func process(head *Node) (int, int) {
	if head == nil {
		return 0, 0
	}
	queue := make([]*Node, 0)
	queue = append(queue, head)
	height, count := 0, 1
	curEnd := head
	var nextEnd *Node
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
			nextEnd = node.Left
			count++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextEnd = node.Right
			count++
		}
		if node == curEnd {
			height++
			curEnd = nextEnd
			nextEnd = nil
		}
	}
	return height, count
}

type Info1 struct {
	Height int
	Count  int
}

func isFull1(head *Node) bool {
	if head == nil {
		return true
	}
	info := process1(head)
	return int(math.Pow(2.0, float64(info.Height))) == info.Count+1
}

func process1(head *Node) *Info1 {
	if head == nil {
		return &Info1{
			Height: 0,
			Count:  0,
		}
	}
	leftInfo := process1(head.Left)
	rightInfo := process1(head.Right)
	height := int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height))) + 1
	count := leftInfo.Count + rightInfo.Count + 1
	return &Info1{Height: height, Count: count}
}

type Info2 struct {
	IsFull bool
	Height int
}

func isFull2(head *Node) bool {
	return process2(head).IsFull
}

func process2(head *Node) *Info2 {
	if head == nil {
		return &Info2{
			IsFull: true,
			Height: 0,
		}
	}
	leftInfo := process2(head.Left)
	rightInfo := process2(head.Right)

	return &Info2{
		IsFull: leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height,
		Height: int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height))) + 1,
	}
}
