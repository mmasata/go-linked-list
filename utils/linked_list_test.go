package utils

import "testing"

func TestPush(t *testing.T) {
	list := LinkedList[string]{}
	Push(&list, "hello")
	Push(&list, "world")

	checkEquals(list.Size, 2, t)
}

func TestPop(t *testing.T) {
	list := LinkedList[string]{}

	Push(&list, "hello")
	Push(&list, "world")
	Push(&list, "test")

	checkEquals(Pop(&list), "test", t)
	checkEquals(Pop(&list), "world", t)
	checkEquals(Pop(&list), "hello", t)
}

func TestSize(t *testing.T) {
	list := LinkedList[string]{}
	checkEquals(list.Size, 0, t)

	Push(&list, "hello")
	Push(&list, "world")
	Push(&list, "test")
	checkEquals(list.Size, 3, t)

	Pop(&list)
	Pop(&list)
	checkEquals(list.Size, 1, t)
}

func TestInsertAfter(t *testing.T) {
	list := LinkedList[string]{}
	Push(&list, "hello")
	InsertAfter(&list, "world", "hello")
	InsertAfter(&list, "test", "world")

	checkEquals(Pop(&list), "test", t)
	checkEquals(Pop(&list), "world", t)
	checkEquals(Pop(&list), "hello", t)
}

func checkEquals[T string | int](got T, expected T, t *testing.T) {
	if got != expected {
		t.Errorf("got %q, wanted %q", got, expected)
	}
}
