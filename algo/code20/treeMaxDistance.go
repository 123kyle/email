package code20

import (
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func maxInstance1(head *Node) int {
	if head == nil {
		return 0
	}
	var arr []*Node
	var getPreList func(head *Node)
	getPreList = func(head *Node) {
		if head == nil {
			return
		}
		arr = append(arr, head)
		getPreList(head.Left)
		getPreList(head.Right)
	}
	getPreList(head)

	var parentMap = make(map[*Node]*Node, len(arr))
	parentMap[head] = nil
	var getParentMap func(head *Node)
	getParentMap = func(head *Node) {
		if head.Left != nil {
			parentMap[head.Left] = head
			getParentMap(head.Left)
		}
		if head.Right != nil {
			parentMap[head.Right] = head
			getParentMap(head.Right)
		}
	}
	getParentMap(head)

	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			max = int(math.Max(float64(max), float64(distance(parentMap, arr[i], arr[j]))))
		}
	}
	return max
}

func distance(parentMap map[*Node]*Node, node1 *Node, node2 *Node) int {
	o1Set := make(map[*Node]struct{})
	cur := node1
	o1Set[cur] = struct{}{}
	_, ok := parentMap[cur]
	for cur != nil && ok {
		cur, ok = parentMap[cur]
		o1Set[cur] = struct{}{}
	}
	cur = node2
	_, ok = o1Set[cur]
	for !ok {
		cur = parentMap[cur]
		_, ok = o1Set[cur]
	}

	lowestAncestor := cur
	cur = node1
	distance1 := 1
	for cur != lowestAncestor {
		cur = parentMap[cur]
		distance1++
	}

	cur = node2
	distance2 := 1
	for cur != lowestAncestor {
		cur = parentMap[cur]
		distance2++
	}
	return distance1 + distance2 - 1
}

type Info struct {
	Height      int
	MaxDistance int
}

func maxDistance2(head *Node) int {
	if head == nil {
		return 0
	}
	return process(head).MaxDistance
}

func process(head *Node) *Info {
	if head == nil {
		return &Info{
			Height:      0,
			MaxDistance: 0,
		}
	}
	leftInfo := process(head.Left)
	rightInfo := process(head.Right)
	var height, maxDistance int
	height = int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height))) + 1

	p1 := leftInfo.MaxDistance
	p2 := rightInfo.MaxDistance
	p3 := leftInfo.Height + rightInfo.Height + 1
	maxDistance = int(math.Max(math.Max(float64(p1), float64(p2)), float64(p3)))

	return &Info{
		Height:      height,
		MaxDistance: maxDistance,
	}
}
