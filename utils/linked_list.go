package utils

type Node[T string | int] struct {
	value T
	next  *Node[T]
}

type LinkedList[T string | int] struct {
	head *Node[T]
	tail *Node[T]
	Size int `default:"0"`
}

func findFirstNodeByValue[T string | int](list *LinkedList[T], value T) *Node[T] {
	var current = list.head
	for current != nil && current.value != value {
		current = current.next
	}

	return current
}

func findNewTail[T string | int](list *LinkedList[T]) *Node[T] {
	var newTail = list.head

	for newTail.next != list.tail {
		newTail = newTail.next
	}

	return newTail
}

func Push[T string | int](list *LinkedList[T], value T) {
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

func InsertAfter[T string | int](list *LinkedList[T], o T, after T) {
	var newNode = &Node[T]{
		value: o,
		next:  nil,
	}

	if list.tail.value == after {
		list.tail.next = newNode
		list.tail = newNode
		list.Size++
	} else {
		var current = findFirstNodeByValue(list, after)
		var oldNext = current.next
		current.next = newNode
		newNode.next = oldNext
		list.Size++
	}
}

func Pop[T string | int](list *LinkedList[T]) T {
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

/*
func PrintList[T string | int](list *LinkedList[T]) {
	var b bytes.Buffer
	b.WriteString("[")

	var current = list.head
	for current != nil {
		if current != list.head {
			b.WriteString(", ")
		}

		b.WriteString(current.value) //TOOD problem with generic
		current = current.next
	}

	b.WriteString("]")
	fmt.Println(b.String())
} */
