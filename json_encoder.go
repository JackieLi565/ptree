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

func (e *JSONEncoder) Encode(tree *Node) string {
	dat, err := json.Marshal(tree)
	if err != nil {
		log.Fatal(err)
	}

	return string(dat)
}

func (e *JSONEncoder) SetRoot(root bool) {
	e.root = root
}
