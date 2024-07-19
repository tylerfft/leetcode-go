package main

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func NewIntensitySegments() IntensitySegments {
	intensitySegments := IntensitySegments{}
	intensitySegments.Tree = redblacktree.NewWithIntComparator()
	return intensitySegments
}

type IntensitySegments struct {
	Tree *redblacktree.Tree
}

func (r *IntensitySegments) Add(from, to, val int) {
	r.addNode(from, val)
	r.addNode(to, -val)
}

func (r *IntensitySegments) Set(from, to, val int) {
	val1, val2 := 0, 0
	keys := r.Tree.Keys()

	for i := range keys {
		key := keys[i].(int)
		node := r.Tree.GetNode(key)
		if key > to {
			break
		}

		val2 += node.Value.(int)
		if key < from {
			val1 += node.Value.(int)
		}

		if key >= from {
			r.Tree.Remove(key)
		}
	}
	r.addNode(from, val-val1)
	r.addNode(to, -val+val2)
}

func (r *IntensitySegments) ToString() string {
	iter := r.Tree.Iterator()
	iter.Begin()

	segs := []string{}
	val := 0
	for iter.Next() {
		node := iter.Node()
		if node.Value.(int) != 0 {
			val += node.Value.(int)
			seg := fmt.Sprintf(`[%v,%v]`, node.Key.(int), val)
			segs = append(segs, seg)
		}
	}
	return "[" + strings.Join(segs, ",") + "]"
}

func (r *IntensitySegments) addNode(key, val int) {
	node := r.Tree.GetNode(key)
	if node == nil {
		r.Tree.Put(key, val)
	} else {
		r.Tree.Put(key, node.Value.(int)+val)
	}
}
