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
	e.SetFancy()
	e.SetIndent("   ")

	t.Scan()
	t.Render(e)
}
