package code12

type Node struct {
	Val  int
	Next *Node
}

func midOrUpMidNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func midOrDownMidNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func midOrUpMidPreNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func midOrDownMidPreNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
