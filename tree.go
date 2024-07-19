package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)


type Tree struct {
	node *Node
	root string
	exclude *Set[string]
	output io.Writer
}

type Node struct {
	Name     string
	Children []*Node
}


type Encoder interface {
	Encode(tree *Node) string
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Render(e Encoder) string {
	result := e.Encode(t.node)

	if t.output != nil {
		t.output.Write([]byte(result))
		t.output.Write([]byte("\n"))
	}

	return result
}

func (t *Tree) SetOutput(output io.Writer) {
	t.output = output
}

func (t *Tree) SetRoot(root string) {
	t.root = root
}

func (t *Tree) SetExclude(exclude *Set[string]) {
	t.exclude = exclude
}

func (t *Tree) Scan() error {
	return filepath.WalkDir(t.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == t.root {
			parts := strings.Split(t.root, string(os.PathSeparator))
			t.node = &Node{
				Name: parts[len(parts) - 1],
			}
			return nil
		}
		relativePath, err := filepath.Rel(t.root, path)
		if err != nil {
			return err
		}

		if t.exclude.Has(d.Name()) {
			return fs.SkipDir
		} 

		parts := strings.Split(relativePath, string(os.PathSeparator))
		current := t.node

		for _, part := range parts {
			found := false
			for _, child := range current.Children {
				if child.Name == part {
					current = child
					found = true
					break
				}
			}

			if !found {
				newNode := &Node{Name: part}
				current.Children = append(current.Children, newNode)
				current = newNode
			}
		}

		return nil
	})
}

