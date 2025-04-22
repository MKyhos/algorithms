package structure

type Node struct {
	Data    any
	IntNode *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Preprend(data any) {
	newNode := &Node{Data: data, IntNode: list.Head}
	list.Head = newNode
	if list.Tail == nil {
		list.Tail = newNode
	}
	list.Size++
}

func (list *LinkedList) Append(data any) {
	newNode := &Node{Data: data, IntNode: nil}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.IntNode = newNode
		list.Tail = newNode
	}
	list.Size++
}

func (list *LinkedList) InsertBefore(data, newData any) bool {
	newNode := &Node{Data: newData}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		list.Size++
		return true
	}

	if list.Head.Data == data {
		newNode.IntNode = list.Head
		list.Head = newNode
		list.Size++
		return true
	}

	// Search
	current := list.Head
	for current != nil && current.IntNode.Data != data {
		current = current.IntNode
	}
	if current != nil {
		newNode.IntNode = current.IntNode
		current.IntNode = newNode
		list.Size++
		return true
	}
	return false
}

func (list *LinkedList) InsertAfter(data, newData any) bool {
	newNode := &Node{Data: newData}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		list.Size++
		return true
	}
	current := list.Head
	for current != nil && current.Data != data {
		current = current.IntNode
	}
	if current != nil {
		newNode = current.IntNode
		current.IntNode = newNode
		list.Size++
		return true
	}
	return false
}

func (list *LinkedList) Delete(data any) bool {
	if list.Head == nil {
		return false // Empty list, nothing to delete
	}

	// check if the head is the node to be deleted
	if list.Head.Data == data {
		list.Head = list.Head.IntNode
		list.Size--
		if list.Head == nil {
			list.Tail = nil
		}
		return true
	}

	current := list.Head
	for current != nil && current.IntNode.Data != data {
		current = current.IntNode
	}
	if current.IntNode != nil {
		current.IntNode = current.IntNode.IntNode
		list.Size--
		return true
	}
	return false
}
