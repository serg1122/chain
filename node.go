package chain

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func CreateNode(value interface{}) *Node {

	return &Node{
		value: value,
	}
}

func (n Node) HasPrev() bool {

	return n.prev != nil
}

func (n *Node) setPrev(prev *Node) {

	n.prev = prev
}

func (n Node) Prev() *Node {

	return n.prev
}

func (n Node) HasNext() bool {

	return n.next != nil
}

func (n *Node) setNext(next *Node) {

	n.next = next
}

func (n *Node) Next() *Node {

	return n.next
}

func (n Node) Value() interface{} {

	return n.value
}

func (n *Node) SetValue(value interface{}) {

	n.value = value
}
