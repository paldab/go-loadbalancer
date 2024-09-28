package queue

import (
	"testing"
)

func setupQueue() *Queue[string] {
	q := NewQueue[string]()

	return &q
}

func TestEnqueue(t *testing.T) {
	q := setupQueue()

	q.Enqueue("test")

	if len(q.Data) != 1 {
		t.Error("Could not add to the queue")
	}

	if q.Data[0] != "test" {
		t.Error("Could not add the correct value to queue")
	}
}

func TestDequeue(t *testing.T) {
	q := setupQueue()

	q.Enqueue("test")

	val, ok := q.Dequeue()

	if !ok || val != "test" {
		t.Errorf("Expected to dequeue 'test', but got %s", val)
	}
}

func TestLen(t *testing.T) {
	q := setupQueue()

	q.Enqueue("test")

	qLen := q.Len()
	if qLen != 1 {
		t.Errorf("Expected queue length of 1, but got %d", qLen)
	}
}

func TestClear(t *testing.T) {
	q := setupQueue()

	q.Enqueue("test")
	q.Enqueue("test1")
	q.Enqueue("test2")
	q.Enqueue("test3")
	q.Enqueue("test4")

	q.Clear()

	qLen := q.Len()
	if qLen != 0 {
		t.Errorf("Expected queue to be empty but got %d items in the queue", qLen)
	}
}
