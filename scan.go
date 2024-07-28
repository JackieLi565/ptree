package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func (t *Tree) Scan() error {
	if t.fullPath {
		return filepath.WalkDir(t.root, t.scanRelativePath)
	} else {
		return filepath.WalkDir(t.root, t.scanAbsolutePath)
	}
}

func (t *Tree) scanAbsolutePath(path string, d fs.DirEntry, err error) error {
	return t.scanPath(path, d, err, func(parts []string) {
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
				newNode := &Node{Name: part, IsDir: d.IsDir()}
				current.Children = append(current.Children, newNode)
				current = newNode
			}
		}
	})
}

func (t *Tree) scanRelativePath(path string, d fs.DirEntry, err error) error {
	return t.scanPath(path, d, err, func(parts []string) {
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
				newNode := &Node{
					Name:  filepath.Join(filepath.Base(t.root), filepath.Join(parts[:i+1]...)),
					IsDir: d.IsDir(),
				}
				current.Children = append(current.Children, newNode)
				current = newNode
			}
		}
	})
}

func (t *Tree) scanPath(path string, d fs.DirEntry, err error, fn func(parts []string)) error {
	if err != nil {
		return err
	}
	if path == t.root {
		t.node = &Node{
			Name:  filepath.Base(t.root),
			IsDir: true,
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
	fn(parts)

	return nil
}
