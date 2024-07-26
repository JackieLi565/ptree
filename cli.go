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

	return root.Execute()
}

func (c *CLI) getEncoder() Encoder {
	switch c.encoding {
	case "json":
		return NewJSONEncoder()
	}

	return NewTextEncoder()
}

func (c *CLI) printStderr(name string, err error) {
	fmt.Printf("Error: %s\n%s\n", name, err)
	os.Exit(1)
}
