package algo

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func maxWidthUseMap(head *Node) int {
	if head == nil {
		return 0
	}
	levelMap := make(map[*Node]int)
	queue := make([]*Node, 0)
	curLevel := 1
	curLevelNodes := 0
	maxWidth := 0
	queue = append(queue, head)
	levelMap[head] = curLevel
	for len(queue) > 0 {
		cur := queue[0]
		curNodeLevel := levelMap[cur]
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			levelMap[cur.Left] = curNodeLevel + 1
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			levelMap[cur.Right] = curNodeLevel + 1
		}
		if curLevel == curNodeLevel {
			curLevelNodes++
		} else {
			maxWidth = max(maxWidth, curLevelNodes)
			curLevel++
			curLevelNodes = 1
		}
	}
	maxWidth = max(maxWidth, curLevelNodes)
	return maxWidth
}

func maxWidthNoMap(head *Node) int {
	if head == nil {
		return 0
	}
	var curEnd, nextEnd *Node
	curEnd = head
	queue := make([]*Node, 0)
	queue = append(queue, head)
	curLevelNodes := 0
	maxWidth := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextEnd = cur.Left
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextEnd = cur.Right
		}
		curLevelNodes++
		if cur == curEnd {
			maxWidth = max(curLevelNodes, maxWidth)
			curLevelNodes = 0
			curEnd = nextEnd
			nextEnd = nil
		}
	}
	return maxWidth
}

