package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Tree map[string]interface{}

func NewTree(root string) (Tree, error) {
	tree := make(Tree)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil // Skip the root directory itself
		}
		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		addToTree(tree, relativePath, d.IsDir())
		return nil
	})
	return tree, err
}

func addToTree(tree Tree, relativePath string, isDir bool) {
	parts := strings.Split(relativePath, string(os.PathSeparator))
	current := tree
	for i, part := range parts {
		if part == "" {
			continue
		}
		if i == len(parts)-1 {
			if isDir {
				current[part] = make(Tree)
			} else {
				current[part] = part
			}
		} else {
			if next, ok := current[part]; ok {
				current = next.(Tree)
			} else {
				next := make(Tree)
				current[part] = next
				current = next
			}
		}
	}
}
