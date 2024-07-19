package main

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet[int]()
	if set.Size() != 0 {
		t.Errorf("expected size 0, got %d", set.Size())
	}
}

func TestAdd(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	if !set.Has(1) {
		t.Errorf("expected set to have 1")
	}
	if set.Size() != 1 {
		t.Errorf("expected size 1, got %d", set.Size())
	}
}

func TestRemove(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Remove(1)
	if set.Has(1) {
		t.Errorf("expected set to not have 1")
	}
	if set.Size() != 0 {
		t.Errorf("expected size 0, got %d", set.Size())
	}
}

func TestHas(t *testing.T) {
	set := NewSet[int]()
	if set.Has(1) {
		t.Errorf("expected set to not have 1")
	}
	set.Add(1)
	if !set.Has(1) {
		t.Errorf("expected set to have 1")
	}
}

func TestSize(t *testing.T) {
	set := NewSet[int]()
	if set.Size() != 0 {
		t.Errorf("expected size 0, got %d", set.Size())
	}
	set.Add(1)
	set.Add(2)
	if set.Size() != 2 {
		t.Errorf("expected size 2, got %d", set.Size())
	}
}

func TestItems(t *testing.T) {
	set := NewSet[int]()
	items := set.Items()
	if len(items) != 0 {
		t.Errorf("expected 0 items, got %d", len(items))
	}
	set.Add(1)
	set.Add(2)
	items = set.Items()
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
	itemMap := make(map[int]bool)
	for _, item := range items {
		itemMap[item] = true
	}
	if !itemMap[1] || !itemMap[2] {
		t.Errorf("expected items to contain 1 and 2")
	}
}
