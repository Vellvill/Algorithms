package double_linked_list

type DLL struct {
	nextNode *DLL
	prevNode *DLL
	data     int
}

func NewDLL(v int) *DLL {
	return &DLL{nil, nil, v}
}

func (n *DLL) AddNode(v int) {
	for i := n; i != nil; i = i.nextNode {
		if i.nextNode == nil {
			new := NewDLL(v)
			n.nextNode = new
			break
		}
	}
}

func (n *DLL) FindMiddle() *DLL {
	for i, j := n, n.FindHead(); i != j; i, j = i.nextNode, j.prevNode {
		if i == j {
			return i
		}
	}
	return nil
}

func (n *DLL) FindHead() *DLL {
	for i := n; i != nil; i = i.nextNode {
		if i.nextNode == nil {
			return i
		}
	}
	return nil
}
