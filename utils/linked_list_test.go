package utils

import "testing"

func TestPush(t *testing.T) {
	list := LinkedList[string]{}
	Push(&list, "hello")
	Push(&list, "world")

	if list.Size != 2 {
		t.Errorf("got %q, wanted %q", list.Size, 2)
	}
}

func TestPop(t *testing.T) {
	var expected = "test"
	list := LinkedList[string]{}

	Push(&list, "hello")
	Push(&list, "world")
	Push(&list, expected)

	var got = Pop(&list)
	if got != expected {
		t.Errorf("got %q, wanted %q", got, expected)
	}
}
