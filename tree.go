package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

func (t *Tree) Scan() error {
	if t.fullPath {
		return filepath.WalkDir(t.root, t.scanRelativePath)
	} else {
		return filepath.WalkDir(t.root, t.scanAbsolutePath)
	}
}

func (t *Tree) scanAbsolutePath(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if path == t.root {
		t.node = &Node{
			Name: filepath.Base(t.root),
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
}

func (t *Tree) scanRelativePath(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if path == t.root {
		t.node = &Node{
			Name: filepath.Base(t.root),
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

	for i := range parts {
		found := false
		for _, child := range current.Children {
			if child.Name == filepath.Join(filepath.Base(t.root), filepath.Join(parts[:i+1]...)) {
				current = child
				found = true
				break
			}
		}

		if !found {
			newNode := &Node{Name: filepath.Join(filepath.Base(t.root), filepath.Join(parts[:i+1]...))}
			current.Children = append(current.Children, newNode)
			current = newNode
		}
	}

	return nil
}
