package queue

import "errors"

type (
	Queue[T any] interface {
		Enqueue(item T)
		Dequeue() (T, error)
		Front() (T, error)
		Size() int
		IsEmpty() bool
	}

	queue[T any] struct {
		items []T
	}
)

func NewQueue[T any]() Queue[T] {
	return &queue[T]{items: make([]T, 0)}
}

func (s *queue[T]) Enqueue(item T) {
	s.items = append(s.items, item)
}

func (s *queue[T]) Dequeue() (T, error) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item, nil
}

func (s *queue[T]) Front() (T, error) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	return s.items[0], nil
}

func (s *queue[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *queue[T]) Size() int {
	return len(s.items)
}
