package utils

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	Size int `default:"0"`
}

func findNewTail[T any](list *LinkedList[T]) *Node[T] {
	var newTail = list.head

	for newTail.next != nil {
		newTail = newTail.next
	}

	return newTail
}

func Push[T any](list *LinkedList[T], value T) {
	var newNode = &Node[T]{
		value: value,
		next:  nil,
	}

	if list.Size == 0 {
		list.head = newNode
		list.tail = newNode
		list.Size = 1
	} else {
		list.tail.next = newNode
		list.tail = newNode
		list.Size++
	}

}

func Pop[T any](list *LinkedList[T]) T {
	var toReturn T
	switch list.Size {
	case 0:
	case 1:
		toReturn = list.head.value
		list.head = nil
		list.tail = nil
		list.Size = 0
	default:
		toReturn = list.tail.value
		var newTail = findNewTail(list)

		list.tail = newTail
		newTail.next = nil
		list.Size--
	}

	return toReturn
}
