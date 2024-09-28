package queue

type Queue[T any] struct {
	Data []T
}

func (q *Queue[T]) Enqueue(item T) {
	if q.Data == nil {
		q.Data = make([]T, 0)
	}

	q.Data = append(q.Data, item)
}

func (q Queue[T]) Len() int {
	if q.Data == nil {
		return 0
	}

	return len(q.Data)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Data == nil || q.Len() == 0 {
		var zero T
		return zero, false
	}

	dequeue := q.Data[0]

	q.Data = q.Data[1:]

	return dequeue, true
}

func (q *Queue[T]) Clear() {
	if q.Len() == 0 {
		return
	}

	hasNext := true
	for hasNext {
		_, ok := q.Dequeue()

		if !ok {
			hasNext = false
		}
	}
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}
