package code04

type Node struct {
	Val  int
	Next *Node
}

func getIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	} else if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

func getLoopNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head.Next.Next, head.Next
	for fast != slow {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func noLoop(head1, head2 *Node) *Node {
	cur1, cur2 := head1, head2
	counter := 0
	for cur1 != nil {
		counter++
		cur1 = cur1.Next
	}
	for cur2 != nil {
		counter--
		cur2 = cur2.Next
	}
	if counter > 0 {
		for i := 0; i < counter; i++ {
			head1 = head1.Next
		}
	} else {
		for i := 0; i < counter; i++ {
			head2 = head2.Next
		}
	}

	for head1 != head2 {
		head1 = head1.Next
		head2 = head2.Next
	}
	return head1
}

func bothLoop(head1, loop1, head2, loop2 *Node) *Node {
	if loop1 == loop2 {
		cur1, cur2 := head1, head2
		counter := 0
		for cur1 != loop1 {
			counter++
			cur1 = cur1.Next
		}
		for cur2 != loop2 {
			counter--
			cur2 = cur2.Next
		}
		if counter > 0 {
			for i := 0; i < counter; i++ {
				head1 = head1.Next
			}
		} else {
			for i := 0; i < counter; i++ {
				head2 = head2.Next
			}
		}

		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur1 := loop1.Next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = loop1.Next
		}
		return nil
	}
}
