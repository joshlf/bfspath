bfspath
=======

A breadth-first-search-based alternative to Dijkstra's algorithm with better runtime in certain cases


Overview
========

Dijkstra's pathfinding algorithm takes the approach of, at every point, prioritizing building up information about the current shortest path.  Thus, since the shortest path is always the one being considered, when the destination is reached, it is guaranteed that the path used to find it is the shortest one.  This is particularly useful on graphs with varying edge-lengths.  However, this requires keeping track of a large amount of information relating to each vertex and path, in data structures with non-constant insertion or removal such as priority queues.

The approach of this algorithm is to mutate graphs with varying edge-lengths into ones in which each edge is considered to be of some unit length.  This allows the algorithm to take advantage of the property of breadth-first-search (BFS) that "levels" of a graph are visited in increasing order.  That is, first vertices which are one edge away from the start, then two, then three, etc.  In a graph in which each edge length is the same, the first path to reach the destination must be the shortest one (and has a distance equal to the number of layers of BFS performed).


Specifics
=========

In order to convert a graph, the distance between two vertices is represented as a counter in terms of some base quantized unit (in my implementation, this is an unsigned integer).  Instead of physically mutating an edge of length 3 into 3 edges of length 1, each pass of BFS simply decrements this counter, and doesn't proceed to the next vertex until the counter reaches 1.

The algorithm maintains two linked lists of vertices, A and B.  At the start, list A simply consists of the start vertex, and B is empty.  Then, in each pass of BFS, each vertex in list A is considered in order.  For each vertex, all adjacent edges directed outward are considered.  If the length is greater than 1, it is decremented, and the vertex is added to list B.  If the length is 1, the vertex at the end of the edge is inspected.  If it has already been visited, it is ignored (because the path corresponding to the previous visit must have a distance no greater than that of the current path).  If it hasn't been visited, its previous pointer is set to point at the current vertex, and it is added to list B.

At the end of the pass of BFS, the contents of list A are replaced by those of list B, and list B is emptied.

Once the destination vertex is found, the algorithm is finished.


Running Time
============

Like normal BFS traversal, BFS path-finding considers each edge once and each vertex once (edges here refers to "virtual" edges - the edges that would exist if a graph were truly to be converted to quantized unit-length edges). In practice, lists A and B are likely implemented as linked-lists or circular buffers or some other constant-time insertion/removal data structure. Thus, the consideration of each vertex is amortized constant time. This gives us a time complexity, for set E of edges and V of vertices, of O(|E| + |V|). Not counting the graph itself, up to |V| vertices must be in a list at any given time, giving a space complexity of O(|V|).

By comparison, Wikipedia lists Dijkstra's worst-case running time as O(|E| + |V| log(|V|)) for a graph with E edges and V vertices.