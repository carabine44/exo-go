package main

import (
	"errors"
	"sync"
)

var ErrDuplicate = errors.New("note already exists")
var ErrNotFound = errors.New("note not found")
var ErrValidation = errors.New("title is required")

type Note struct {
	ID      string
	Title   string
	Content string
}

type NoteStore interface {
	Save(n *Note) error
	Get(id string) (*Note, error)
	All() []*Note
}

type MemoryStore struct {
	mu    sync.Mutex
	notes map[string]*Note
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		notes: make(map[string]*Note),
	}
}

func (m *MemoryStore) Save(n *Note) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if n.Title == "" {
		return ErrValidation
	}

	if _, exists := m.notes[n.ID]; exists {
		return ErrDuplicate
	}

	m.notes[n.ID] = n

	return nil
}

func (m *MemoryStore) Get(id string) (*Note, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	note, exists := m.notes[id]

	if !exists {
		return nil, ErrNotFound
	}

	return note, nil
}

func (m *MemoryStore) All() []*Note {
	m.mu.Lock()
	defer m.mu.Unlock()

	var result []*Note

	for _, note := range m.notes {
		result = append(result, note)
	}

	return result
}
