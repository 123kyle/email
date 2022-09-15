package code03

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{Val: root.Val}
	head.Left = en(root.Children)
	return head
}

func en(children []*Node) *TreeNode {
	var head, cur *TreeNode

	for _, child := range children {
		t := &TreeNode{Val: child.Val}
		if head == nil {
			head = t
		} else {
			cur.Right = t
		}
		cur = t
		cur.Left = en(child.Children)
	}
	return head
}

func decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}
	return &Node{root.Val, de(root.Left)}
}

func de(root *TreeNode) []*Node {
	var children []*Node
	for root != nil {
		cur := &Node{root.Val, de(root.Left)}
		children = append(children, cur)
		root = root.Right
	}
	return children
}
