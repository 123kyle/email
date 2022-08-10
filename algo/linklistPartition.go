package algo


type ListNode struct {
	Val int
	Next *ListNode
}


func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var sH *ListNode
	var sT *ListNode
	var eH *ListNode
	var eT *ListNode
	var mH *ListNode
	var mT *ListNode

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		if cur.Val < x {
			if sH == nil {
				sH = cur
				sT = cur
			} else {
				sT.Next = cur
				sT = cur
			}
		} else if cur.Val > x {
			if mH == nil {
				mH = cur
				mT = cur
			} else {
				mT.Next = cur
				mT = cur
			}
		} else {
			if eH == nil {
				eH = cur
				eT = cur
			} else {
				eT.Next = cur
				eT = cur
			}
		}
		cur = next
	}

	if sH != nil {
		sT.Next = eH
	}

	if eH != nil {
		eT.Next = mH
	}
	if sH != nil {
		return sH
	}
	if eH != nil {
		return eH
	}
	return mH
}
