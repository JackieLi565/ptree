package main

import "os"

func main() {
	cwd, _ := os.Getwd()
	s := NewSet[string]()
	s.Add(".git")

	t := NewTree()
	t.SetRoot(cwd)
	t.SetExclude(s)
	t.SetOutput(os.Stdout)

	e := NewTextEncoder()

	t.Scan()
	t.Render(e)
}
