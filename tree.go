package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Tree struct {
	node     *Node
	root     string
	fullPath bool
	exclude  *Set[string]
	output   io.Writer
}

type Node struct {
	Name     string
	Children []*Node
	IsDir    bool
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

func (t *Tree) SetRoot(root string) error {
	absPath, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	stat, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory not found %w", err)
	}
	if !stat.IsDir() {
		return fmt.Errorf("path is not a directory: %s", root)
	}

	t.root = absPath
	return nil
}

func (t *Tree) SetOutput(output io.Writer) {
	t.output = output
}

func (t *Tree) SetExclude(exclude *Set[string]) {
	t.exclude = exclude
}
