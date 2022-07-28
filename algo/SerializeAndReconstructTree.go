package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func preSerial(head *Node) []string {
	ans := make([]string, 0)
	var pres func(head *Node)
	pres = func(head *Node) {
		if head == nil {
			ans = append(ans, "nil")
		} else {
			ans = append(ans, strconv.Itoa(head.Val))
			pres(head.Left)
			pres(head.Right)
		}
	}
	pres(head)
	return ans
}

func inSerial(head *Node) []string {
	ans := make([]string, 0)
	var ins func(head *Node)
	ins = func(head *Node) {
		if head == nil {
			ans = append(ans, "nil")
		} else {
			ins(head.Left)
			ans = append(ans, strconv.Itoa(head.Val))
			ins(head.Right)
		}
	}
	ins(head)
	return ans
}

func posSerial(head *Node) []string {
	ans := make([]string, 0)
	var poss func(head *Node)
	poss = func(head *Node) {
		if head == nil {
			ans = append(ans, "nil")
		} else {
			poss(head.Left)
			poss(head.Right)
			ans = append(ans, strconv.Itoa(head.Val))
		}
	}
	poss(head)
	return ans
}

func buildByPreQueue(preList []string) *Node {
	if preList == nil || len(preList) == 0 {
		return nil
	}
	var preb func([]string) *Node
	preb = func(strings []string) *Node {
		v := preList[0]
		preList = preList[1:]
		if v == "nil" {
			return nil
		}
		head := new(Node)
		head.Val, _ = strconv.Atoi(v)
		head.Left = preb(preList)
		head.Right = preb(preList)
		return head
	}

	return preb(preList)
}

func buildByPosQueue(posList []string) *Node {
	if posList == nil || len(posList) == 0 {
		return nil
	}
	var posb func([]string) *Node

	posb = func(strings []string) *Node {
		length := len(posList)
		v := posList[length-1]
		posList = posList[:length-1]
		if v == "nil" {
			return nil
		}
		head := new(Node)
		head.Val, _ = strconv.Atoi(v)
		head.Right = posb(posList)
		head.Left = posb(posList)
		return head
	}

	return posb(posList)
}

func levelSerial(head *Node) []string {
	ans := make([]string, 0)
	if head == nil {
		ans = append(ans, "nil")
	} else {
		ans = append(ans, strconv.Itoa(head.Val))
		queue := make([]*Node, 0)
		queue = append(queue, head)
		for len(queue) > 0 {
			head = queue[0]
			queue = queue[1:]
			if head.Left != nil {
				queue = append(queue, head.Left)
				ans = append(ans, strconv.Itoa(head.Left.Val))
			} else {
				ans = append(ans, "nil")
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
				ans = append(ans, strconv.Itoa(head.Right.Val))
			} else {
				ans = append(ans, "nil")
			}
		}
	}
	return ans
}

func buildByLevelQueue(levelList []string) *Node {
	if levelList == nil || len(levelList) == 0 {
		return nil
	}
	v := levelList[0]
	levelList = levelList[1:]
	head := generateNode(v)
	queue := make([]*Node, 0)
	if head != nil {
		queue = append(queue, head)
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left = generateNode(levelList[0])
		levelList = levelList[1:]
		node.Right = generateNode(levelList[0])
		levelList = levelList[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return head
}

func generateNode(val string) *Node {
	if val == "nil" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	return &Node{Val: v}
}

func generateRandomBST(maxLevel, maxValue int) *Node {
	return generate(1, maxLevel, maxValue)
}

func generate(level int, maxLevel int, maxValue int) *Node {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &Node{Val: rand.Intn(maxValue) + 1}
	head.Left = generate(level+1, maxLevel, maxValue)
	head.Right = generate(level+1, maxLevel, maxValue)
	return head
}

func isSameValueStructure(head1, head2 *Node) bool {
	if (head1 == nil && head2 != nil) || (head2 == nil && head1 != nil) {
		return false
	}
	if head1 == nil && head2 == nil {
		return true
	}
	if head1.Val != head2.Val {
		return false
	}
	return isSameValueStructure(head1.Left, head2.Left) && isSameValueStructure(head1.Right, head2.Right)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxLevel := 5
	maxVal := 100
	testTimes := 1000000
	for i := 0; i < testTimes; i++ {
		head := generateRandomBST(maxLevel, maxVal)
		pre := preSerial(head)
		pos := posSerial(head)
		level := levelSerial(head)
		preBuild := buildByPreQueue(pre)
		posBuild := buildByPosQueue(pos)
		levelBuild := buildByLevelQueue(level)
		if !isSameValueStructure(preBuild, posBuild) || !isSameValueStructure(preBuild, levelBuild) {
			fmt.Println("Oops")
			return
		}
	}
	fmt.Println("Nice")
}
