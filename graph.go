package graph

type graphNode struct {
	// List of adjacent nodes and distances
	next []struct {
		dist uint64
		node *graphNode
	}
	previous *graphNode
}

type llNode struct {
	node *graphNode
	next *llNode
}

type ll struct {
	head *llNode
	tail *llNode
}

func (l *ll) add(n *graphNode) {
	newNode := new(llNode)
	newNode.node = n
	
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

// Takes start and end nodes, and returns path length
func bfsPath(start, end *graphNode) int {
	// Set start's previous pointer
	// because nil is a sentinal value.
	// This is never actually used, so 
	// it doesn't cause recursion problems
	start.previous = start
	
	listA := new(ll)
	listB := new(ll)
	listA.add(start)
	
	// Each iteration of BFS increases
	// path length by 1
	i := 1
	for ; ; i++ {
		for cur := listA.head; cur != nil ; cur = cur.next {
			
			// Iterate through this node's adjacent nodes
			for j, adj := range cur.node.next {
				if adj.dist < 2 && adj.node.previous == nil {
					adj.node.previous = cur.node
					if adj.node == end {
						break
					}
					listB.add(adj.node)
				} else {
					cur.node.next[j].dist--
					
					// If this node hasn't been added to listB yet
					if listB.head == nil || listB.tail.node != cur.node {
						listB.add(cur.node)
					}
				}
			}
		}
		if end.previous != nil {
			break
		}
		listA = listB
		listB = new(ll)
	}

	// Undo the change from
	// the beginning
	start.previous = nil
	return i
}