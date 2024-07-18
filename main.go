package main

import (
	"fmt"
)

func main() {
	test := "/Users/jackieli/github/JackieLi565/ptree"

	tree, err := NewTree(test)
	fmt.Println(err)
	printTree(tree, 0)

}

func printTree(tree Tree, level int) {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}
	for name, value := range tree {
		fmt.Println(prefix + name)
		if subTree, ok := value.(Tree); ok {
			printTree(subTree, level+1)
		}
	}
}
