package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CLI struct {
	tree       *Tree
	silent     bool
	encoding   string
	exclude    []string
	outputFile string
	fancy      bool
	root       bool
	trailing   bool
}

func NewCLI() *CLI {
	tree := NewTree()
	tree.SetExclude(NewSet[string]())

	return &CLI{
		tree: tree,
	}
}

func (c *CLI) Run() error {

	root := &cobra.Command{
		Use:   "ptree [directory]",
		Short: "Generate project structures via a CLI",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := c.tree.SetRoot(args[0]); err != nil {
				c.printStderr("Directory lookup", err)
				return
			}

			for _, file := range c.exclude {
				c.tree.exclude.Add(file)
			}

			if !c.silent {
				c.tree.SetOutput(os.Stdout)
			}

			if err := c.tree.Scan(); err != nil {
				c.printStderr("Directory scan", err)
			}

			result := c.tree.Render(c.getEncoder())
			if c.outputFile != "" {
				if err := os.WriteFile(c.outputFile, []byte(result), 0644); err != nil {
					c.printStderr("File output", err)
				}
			}
		},
	}

	// short forms
	root.Flags().BoolVarP(&c.tree.fullPath, "full-path", "f", false, "Show full path of project tree")
	root.Flags().BoolVarP(&c.silent, "silent", "s", false, "Suppress output to stdout")
	root.Flags().StringVarP(&c.encoding, "encoding", "e", "text", "Project tree encoding")

	// long forms
	root.Flags().StringSliceVar(&c.exclude, "exclude", make([]string, 0), "Exclude directories or files from project tree")
	root.Flags().StringVar(&c.outputFile, "output-file", "", "Project tree output file")
	root.Flags().BoolVar(&c.fancy, "fancy", false, "Formatted fancy output")
	root.Flags().BoolVar(&c.root, "root", false, "Include root directory in final output")
	root.Flags().BoolVar(&c.trailing, "trailing", false, "Include trailing slash")

	return root.Execute()
}

func (c *CLI) getEncoder() Encoder {
	switch c.encoding {
	case "json":
		e := NewJSONEncoder()
		e.SetRoot(c.root)

		return e
	case "text":
		e := NewTextEncoder()
		e.SetFancy(c.fancy)
		e.SetRoot(c.root)
		e.SetTrailing(c.trailing)

		return e
	}

	c.printWarn("Invalid 'encoding' parameter", "defaulting to a standard text encoding")
	return NewTextEncoder()
}

func (c *CLI) printStderr(name string, err error) {
	fmt.Printf("Error: %s\n%s\n", name, err)
	os.Exit(1)
}

func (c *CLI) printWarn(name string, description string) {
	fmt.Printf("Warning: %s\n%s\n", name, description)
}
