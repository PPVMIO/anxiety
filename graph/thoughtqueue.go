package graph

import "sync"

// ThoughtQueue the queue of Thoughts
type ThoughtQueue struct {
	items []Thought
	lock  sync.RWMutex
}

// New creates a new ThoughtQueue
func (s *ThoughtQueue) New() *ThoughtQueue {
	s.lock.Lock()
	s.items = []Thought{}
	s.lock.Unlock()
	return s
}

// Enqueue adds an Thought to the end of the queue
func (s *ThoughtQueue) Enqueue(t Thought) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Dequeue removes an Thought from the start of the queue
func (s *ThoughtQueue) Dequeue() *Thought {
	s.lock.Lock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	s.lock.Unlock()
	return &item
}

// Front returns the item next in the queue, without removing it
func (s *ThoughtQueue) Front() *Thought {
	s.lock.RLock()
	item := s.items[0]
	s.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (s *ThoughtQueue) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}

// Size returns the number of Thoughts in the queue
func (s *ThoughtQueue) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
