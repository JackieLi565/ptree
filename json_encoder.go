package main

import (
	"encoding/json"
	"log"
)

type JSONEncoder struct {
	root bool
}

func NewJSONEncoder() *JSONEncoder {
	return &JSONEncoder{}
}

func (w *JSONEncoder) Encode(tree *Node) string {
	dat, err := json.Marshal(tree)
	if err != nil {
		log.Fatal(err)
	}

	return string(dat)
}

func (w *JSONEncoder) SetRoot() bool {
	w.root = !w.root
	return w.root
}
