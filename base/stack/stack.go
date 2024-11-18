package stack

import "errors"

type (
	Stack[T any] interface {
		Push(item T)
		Pop() (T, error)
		Peek() (T, error)
		Size() int
		IsEmpty() bool
	}

	stack[T any] struct {
		items []T
	}
)

func NewStack[T any]() Stack[T] {
	return &stack[T]{items: make([]T, 0)}
}

func (s *stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack[T]) Size() int {
	return len(s.items)
}
