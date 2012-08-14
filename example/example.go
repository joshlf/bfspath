package main

import (
	"fmt"
	"bfspath"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <filename>", os.Args[0])
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error openning file: %v\n", err)
		os.Exit(2)
	}
	start, end := importGraph(file)
	fmt.Printf("Shortest path has length %v\n", graph.FindPath(start))
	for cur := graph.Node(end); !cur.IsStart(); cur = cur.Prev() {
		c, _ := cur.(*node)
		fmt.Printf("%v <- ", c.name)
	}
	fmt.Printf("%v\n", start.name)
}
