// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	INERR = 3
	INLEN = 4
	MFD   = 5
)

var nodes = make(map[string]*node)

// Reads a graph definition, constructs
// the graph, and then returns the start
// node.
func importGraph(rdr io.Reader) (start, end *node) {
	r := bufio.NewReader(rdr)

	// Handle first line
	b, ip, err := r.ReadLine()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(INERR)
	} else if ip {
		fmt.Println("Input too long.")
		os.Exit(INLEN)
	}
	strs := strings.Fields(string(b))
	if len(strs) != 2 {
		fmt.Println("Malformed definition: expected <start> <end> on first line")
		os.Exit(MFD)
	}
	if strs[0] == strs[1] {
		nodes[strs[0]] = new(node)
		n := nodes[strs[0]]
		n.isStart = true
		n.isEnd = true
		n.name = strs[0]
		start = n
		end = n
	} else {
		nodes[strs[0]] = new(node)
		nodes[strs[1]] = new(node)
		n := nodes[strs[0]]
		m := nodes[strs[1]]
		n.name = strs[0]
		m.name = strs[1]
		n.isStart = true
		m.isEnd = true
		start = n
		end = m
	}

	b, ip, err = r.ReadLine()
	for ; err == nil || err.Error() != "EOF"; b, ip, err = r.ReadLine() {
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			os.Exit(INERR)
		} else if ip {
			fmt.Println("Input too long.")
			os.Exit(INLEN)
		}
		strs := strings.Fields(string(b))
		if len(strs) != 3 {
			fmt.Println("Malformed definition: expected 3 words per line")
			os.Exit(MFD)
		}
		length, err := strconv.Atoi(strs[2])
		if err != nil {
			fmt.Println("Malformed definition: expected edge length as third item on line")
			os.Exit(MFD)
		}
		n, ok := nodes[strs[0]]
		if !ok {
			nodes[strs[0]] = new(node)
			n = nodes[strs[0]]
			n.name = strs[0]
		}
		m, ok := nodes[strs[1]]
		if !ok {
			nodes[strs[1]] = new(node)
			m = nodes[strs[1]]
			m.name = strs[1]
		}
		n.addEdge(m, length)
	}
	end.isEnd = true
	return
}

func (n *node) addEdge(m *node, dist int) {
	n.next = append(n.next, next{uint64(dist), uint64(dist), m})
}
