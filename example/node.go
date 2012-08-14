package main

import "bfspath"

type next struct {
	// dist remains constant
	dist uint64
	// counter for graph.FindPath to decrement
	counter uint64
	node    graph.Node
}

type node struct {
	// List of adjacent nodes and distances
	next     []next
	previous graph.Node
	isEnd    bool
	isStart  bool

	// A convenience so we can print the
	// path in a human-readable format
	name string
}

func (n *node) Edges() int            { return len(n.next) }
func (n *node) Next(i int) graph.Node { return n.next[i].node }
func (n *node) Dec(i int) bool {
	if n.next[i].counter > 1 {
		n.next[i].counter--
		return false
	}
	return true
}
func (n *node) PrevCheck() bool {
	return n.previous != nil
}
func (n *node) PrevSet(m graph.Node) {
	n.previous = m
}
func (n *node) Prev() graph.Node { return n.previous }
func (n *node) IsEnd() bool      { return n.isEnd }
func (n *node) IsStart() bool    { return n.isStart }
