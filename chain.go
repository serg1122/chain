package chain

type Chain struct {
	first *Node
	last  *Node
}

func New() *Chain {
	return &Chain{}
}

func (chain *Chain) IsEmpty() bool {

	return chain.first == nil
}

func (chain *Chain) init(node *Node) {

	node.setPrev(nil)
	node.setNext(nil)
	chain.first = node
	chain.last = node
}

func (chain *Chain) Prepend(node *Node) {

	if chain.IsEmpty() {

		chain.init(node)
		return
	}

	node.setNext(chain.first)
	chain.first.setPrev(node)
	chain.first = node
}

func (chain *Chain) Append(node *Node) {

	if chain.IsEmpty() {

		chain.init(node)
		return
	}

	node.setPrev(chain.last)
	chain.last.setNext(node)
	chain.last = node
}

func (chain Chain) First() *Node {

	if chain.IsEmpty() {

		return nil
	}

	return chain.first
}

func (chain Chain) Last() *Node {

	if chain.IsEmpty() {

		return nil
	}

	return chain.last
}
