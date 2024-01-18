package mpiserver

import (
	"errors"
	"fmt"
	"go-distributed/representations"
	"sync"
)

// SafeStack represents a thread-safe stack data structure.
type SafeStack struct {
	items []representations.Representation
	mu    sync.Mutex
}

// Push adds an item to the top of the stack.
func (s *SafeStack) Push(item representations.Representation) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// It returns an error if the stack is empty.
func (s *SafeStack) Pop() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index] // Remove the top element
	return item, nil
}

// Peek returns the item from the top of the stack without removing it.
// It returns an error if the stack is empty.
func (s *SafeStack) Peek() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *SafeStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *SafeStack) Size() int {
	return len(s.items)
}

func TestRepository() {
	// Create a new thread-safe stack
	stack := SafeStack{}

	// Use the stack concurrently from multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			stack.Push(&representations.BinaryRepresentation{
				Genes: []int{1, 0, 1, 0},
			})
			topItem, err := stack.Peek()
			if err == nil {
				fmt.Printf("Goroutine %d: Top item: %v\n", index, topItem)
			}

			item, err := stack.Pop()
			if err == nil {
				fmt.Printf("Goroutine %d: Popped: %v\n", index, item)
			}
		}(i)
	}

	wg.Wait()
}
