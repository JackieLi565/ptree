package main

import "strings"

type TextEncoder struct {
	root     bool
	fancy    bool
	trailing bool
}

type CharSet struct {
	tail   string
	branch string
	bar    string
}

const (
	Root   = "."
	Indent = "    "

	UTF8Tail   = "└── "
	UTF8Branch = "├── "
	UTF8Bar    = "│    "

	ASCIITail   = "`-- "
	ASCIIBranch = "|-- "
	ASCIIBar    = "|    "
)

func NewTextEncoder() *TextEncoder {
	return &TextEncoder{}
}

func (e *TextEncoder) Encode(tree *Node) string {
	var sb strings.Builder

	if e.root {
		return sb.String()
	} else {
		if e.fancy {
			e.fancyEncode(tree, &sb)
			return sb.String()
		} else {
			e.normalEncode(tree, &sb)
			return sb.String()
		}
	}
}

func (e *TextEncoder) SetRoot() bool {
	e.root = !e.root
	return e.root
}

func (e *TextEncoder) SetFancy() bool {
	e.fancy = !e.fancy
	return e.fancy
}

func (e *TextEncoder) SetTrailing() bool {
	e.trailing = !e.trailing
	return e.trailing
}

func (e *TextEncoder) fancyEncode(node *Node, sb *strings.Builder) {
	charSet := CharSet{
		bar:    UTF8Bar,
		branch: UTF8Branch,
		tail:   UTF8Tail,
	}

	e.encode(node, sb, "", true, &charSet)
}

func (e *TextEncoder) normalEncode(node *Node, sb *strings.Builder) {
	charSet := CharSet{
		bar:    ASCIIBar,
		branch: ASCIIBranch,
		tail:   ASCIITail,
	}

	e.encode(node, sb, "", true, &charSet)
}

func (e *TextEncoder) encode(node *Node, sb *strings.Builder, prefix string, isLast bool, charSet *CharSet) {
	sb.WriteString(prefix)
	if isLast {
		sb.WriteString(charSet.tail)
		prefix += Indent
	} else {
		sb.WriteString(charSet.branch)
		prefix += charSet.bar
	}
	sb.WriteString(node.Name + "\n")

	for i, child := range node.Children {
		e.encode(child, sb, prefix, i == len(node.Children)-1, charSet)
	}
}
