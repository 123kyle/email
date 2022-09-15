package code22

type Interface interface {
	Insert(word string)
	Delete(word string)
	Search(word string) int
	PrefixCount(word string) int
}

type Node struct {
	Pass      int
	End       int
	NextNodes []*Node
}

func NewNode() *Node {
	return &Node{
		NextNodes: make([]*Node, 26),
	}
}

type TrieTree struct {
	Root *Node
}

func NewTrieTree() *TrieTree {
	return &TrieTree{Root: NewNode()}
}

func (t *TrieTree) Insert(word string) {
	curNode := t.Root
	curNode.Pass++
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if curNode.NextNodes[path] == nil {
			curNode.NextNodes[path] = NewNode()
		}
		curNode = curNode.NextNodes[path]
		curNode.Pass++
	}
	curNode.End++
}

func (t *TrieTree) Delete(word string) {
	if t.Search(word) != 0 {
		curNode := t.Root
		curNode.Pass--
		for i := 0; i < len(word); i++ {
			path := word[i] - 'a'
			curNode.NextNodes[path].Pass--
			if curNode.NextNodes[path].Pass == 0 {
				curNode.NextNodes[path] = nil
				return
			}
			curNode = curNode.NextNodes[path]
		}
		curNode.End--
	}
}

func (t *TrieTree) Search(word string) int {
	curNode := t.Root
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if curNode.NextNodes[path] == nil {
			return 0
		}
		curNode = curNode.NextNodes[path]
	}
	return curNode.Pass
}

func (t *TrieTree) PrefixCount(prefix string) int {
	curNode := t.Root
	for i := 0; i < len(prefix); i++ {
		path := prefix[i] - 'a'
		if curNode.NextNodes[path] == nil {
			return 0
		}
		curNode = curNode.NextNodes[path]
	}
	return curNode.Pass
}
