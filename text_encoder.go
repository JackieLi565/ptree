package main

import "strings"

type TextEncoder struct {
	fullPath bool
	root     bool
	fancy    bool
	trailing bool
	indent 	 string
}

type UTF8 string

const (
	UTF8Root UTF8 = "."
	UTF8Tail = "└── "
	UTF8Branch = "├── "
	UTF8Bar = "│    "
)

type ASCII string

const (
	ASCIIRoot ASCII = "."
	ASCIITail = "`--"
	ASCIIBranch = "|--"
	ASCIIBar = "|"
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
			e.fancyEncode(tree, &sb, "", true)
			return sb.String()
		} else {
			e.normalEncode(tree, &sb)
			return sb.String()
		}
	}
}

func (e *TextEncoder) SetFullPath() bool {
	e.fullPath = !e.fullPath
	return e.fullPath
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

func (e *TextEncoder) SetIndent(indent string) {
	e.indent = indent
}

func (e *TextEncoder) fancyEncode(node *Node, sb *strings.Builder, prefix string, isLast bool) {
	sb.WriteString(prefix)
	if isLast {
		sb.WriteString(UTF8Tail)
		prefix += e.indent
	} else {
		sb.WriteString(UTF8Branch)
		prefix += UTF8Bar
	}
	sb.WriteString(node.Name + "\n")

	for i, child := range node.Children {
		e.fancyEncode(child, sb, prefix, i == len(node.Children)-1)
	}
}

func (e *TextEncoder) normalEncode(tree *Node, sb *strings.Builder) {
	
}

func (e *TextEncoder) prefix(sb *strings.Builder, level uint) {
	sb.WriteString(strings.Repeat(e.indent, int(level)))
}
