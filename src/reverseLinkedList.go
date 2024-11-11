package linkedList

type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	Node *Node
}

func ReverseLinkedList(list *LinkedList) *LinkedList {
	current := list.Node
	var previous *Node
	for current != nil {
		new := current.Next
		current.Next = previous
		previous = current
		current = new
	}
	return &LinkedList{Node: previous}
}
