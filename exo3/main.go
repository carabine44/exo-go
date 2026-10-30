package main

import (
	"bufio"
	"os"
	"strings"
)

type Note struct {
	Title   string
	Content string
	Tags    []string
}

func NewNote(title, content string) *Note {
	return &Note{
		Title:   title,
		Content: content,
		Tags:    []string{},
	}
}

func (n *Note) Preview() string {
	if len(n.Content) <= 80 {
		return n.Content
	}

	return n.Content[:80]
}

func (n *Note) AddTag(tag string) {
	for _, t := range n.Tags {
		if t == tag {
			return
		}
	}

	n.Tags = append(n.Tags, tag)
}

func LoadFromFile(path string) ([]*Note, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var notes []*Note

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		note := NewNote("Imported", line)

		notes = append(notes, note)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func HasGoTag(n *Note) bool {
	for _, tag := range n.Tags {
		if strings.ToLower(tag) == "go" {
			return true
		}
	}

	return false
}
