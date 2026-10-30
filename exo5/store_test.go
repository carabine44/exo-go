package main

import (
	"errors"
	"testing"
)

func TestSave_valid(t *testing.T) {
	store := NewMemoryStore()

	note := &Note{
		ID:    "1",
		Title: "Ma note",
	}

	err := store.Save(note)

	if err != nil {
		t.Fatalf("erreur inattendue : %v", err)
	}
}

func TestSave_emptyTitle(t *testing.T) {
	store := NewMemoryStore()

	note := &Note{
		ID:    "1",
		Title: "",
	}

	err := store.Save(note)

	if err == nil {
		t.Fatal("une erreur était attendue")
	}
}

func TestSave_duplicate(t *testing.T) {
	store := NewMemoryStore()

	note := &Note{
		ID:    "1",
		Title: "Test",
	}

	_ = store.Save(note)

	err := store.Save(note)

	if !errors.Is(err, ErrDuplicate) {
		t.Fatal("ErrDuplicate attendu")
	}
}

func TestGet_notFound(t *testing.T) {
	store := NewMemoryStore()

	_, err := store.Get("?")

	if !errors.Is(err, ErrNotFound) {
		t.Fatal("ErrNotFound attendu")
	}
}
