package utils

import (
	"sync"
	"testing"
)

func TestPush(t *testing.T) {
	list := LinkedList[string]{}
	Push(&list, "hello")
	Push(&list, "world")

	checkEqualsInt(list.Size, 2, t)
}

func TestPop(t *testing.T) {
	list := LinkedList[string]{}

	Push(&list, "hello")
	Push(&list, "world")
	Push(&list, "test")

	checkEqualsStr(Pop(&list), "test", t)
	checkEqualsStr(Pop(&list), "world", t)
	checkEqualsStr(Pop(&list), "hello", t)
}

func TestSize(t *testing.T) {
	list := LinkedList[string]{}
	checkEqualsInt(list.Size, 0, t)

	Push(&list, "hello")
	Push(&list, "world")
	Push(&list, "test")
	checkEqualsInt(list.Size, 3, t)

	Pop(&list)
	Pop(&list)
	checkEqualsInt(list.Size, 1, t)
}

func TestInsertAfter(t *testing.T) {
	list := LinkedList[string]{}
	Push(&list, "hello")
	InsertAfter(&list, "world", "hello")
	InsertAfter(&list, "test", "world")

	checkEqualsStr(Pop(&list), "test", t)
	checkEqualsStr(Pop(&list), "world", t)
	checkEqualsStr(Pop(&list), "hello", t)
}

func TestConcurrencyPush(t *testing.T) {
	var iterations = 1000
	var wg sync.WaitGroup
	list := LinkedList[string]{}

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Push(&list, "value")
		}()
	}

	wg.Wait()
	checkEqualsInt(list.Size, iterations, t)
}

func TestConcurrencyPop(t *testing.T) {
	var iterations = 1000
	var wg sync.WaitGroup
	list := LinkedList[string]{}

	//add 1000 value with one thread
	for i := 0; i < iterations; i++ {
		Push(&list, "value")
	}

	//pop 1000 value via goroutines
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Pop(&list)
		}()
	}

	wg.Wait()
	checkEqualsInt(list.Size, 0, t)
}

func checkEqualsStr(got string, expected string, t *testing.T) {
	if got != expected {
		t.Errorf("got %s, wanted %s", got, expected)
	}
}

func checkEqualsInt(got int, expected int, t *testing.T) {
	if got != expected {
		t.Errorf("got %d, wanted %d", got, expected)
	}
}
