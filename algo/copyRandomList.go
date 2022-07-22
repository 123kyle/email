package algo

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 测试链接 : https://leetcode.com/problems/copy-list-with-random-pointer/

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		m[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		node := m[cur]
		node.Next = m[cur.Next]
		node.Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{
			Val:  cur.Val,
			Next: next,
		}
		cur = next
	}
	cur = head
	for cur != nil {
		next := cur.Next.Next
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = next
	}
	cur = head
	newHead := head.Next
	for cur != nil {
		copyCur := cur.Next
		next := cur.Next.Next
		if next != nil {
			copyCur.Next = next.Next
		}
		cur.Next = next
		cur = next
	}

	return newHead
}
