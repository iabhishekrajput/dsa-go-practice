# Graphs — Revision Notes

## How It Works

A **graph** is a set of **nodes** (vertices) connected by **edges**. Unlike trees, graphs have no root, may contain cycles, and may be disconnected.

### Four defining properties
- **Directed** (one-way) vs **undirected** (two-way)
- **Weighted** (edges have values) vs **unweighted**
- **Cyclic** vs **acyclic** (DAGs enable topological sort)
- **Connected** (every node reachable) vs **disconnected** (multiple components)

### Representations

**Adjacency list** (default) — `map[int][]int` or `[][]int`. Space O(V+E). Neighbors of X: O(degree(X)).

**Adjacency matrix** — 2D grid, `matrix[i][j] = 1` if edge. Space O(V²). Edge lookup: O(1). Wasteful for sparse graphs.

**Edge list** — list of `(a, b)` pairs. Rare; useful for edge-iterating algorithms (Kruskal).

### The key difference from trees

Graphs can have cycles and multiple paths — so traversal **requires a `visited` set**. Without it, cycles cause infinite loops.

## Complexity Cheat Sheet

| Operation | Adjacency List | Adjacency Matrix |
|-----------|---------------|------------------|
| Storage | O(V + E) | O(V²) |
| Full BFS/DFS | O(V + E) | O(V²) |
| Check edge (A,B) | O(degree(A)) | O(1) |
| List A's neighbors | O(degree(A)) | O(V) |

## Key Patterns

### Build Graph from Edge List
- **When to use:** Input is `n` nodes + `[][]int` edge pairs
- **How it works:** `map[int][]int`, for each `[a,b]`: append `b` to `adj[a]` and `a` to `adj[b]` (undirected adds both directions)
- **Watch for:** Directed graphs only add one direction. Self-loops and duplicate edges — depends on problem constraints

### Count Connected Components
- **When to use:** How many "islands" in the graph
- **How it works:** Outer loop over all nodes. If unvisited, start BFS/DFS from it (increment counter). Each traversal exhausts one component
- **Watch for:** Isolated nodes (no edges) are still components — don't skip them

### Grid as Implicit Graph
- **When to use:** 2D grid problems with 4-directional connectivity (islands, flood fill, rotting oranges, word search)
- **How it works:** Each cell is a node, neighbors are up/down/left/right. No explicit adjacency list — compute neighbors with coordinate math
- **Watch for:** Bounds checking before indexing. Use a `directions` table `{{-1,0},{1,0},{0,-1},{0,1}}` to iterate cleanly

### Graph Clone (DFS + Memoization)
- **When to use:** Deep copy a graph
- **How it works:** `map[*Node]*Node` from original → clone. For each node: check map first (return existing clone for cycle cases), allocate new clone, register in map **before recursing**, then recurse on neighbors
- **Watch for:** Order matters — **allocate and register in map BEFORE recursing** into neighbors. Otherwise cycles → infinite recursion

### BFS Visited-on-Enqueue
- **When to use:** Any BFS on graphs
- **How it works:** Mark a node visited when you *add it to the queue*, not when you *dequeue it*
- **Watch for:** If you mark on dequeue, the same node can be enqueued many times (once per discovering neighbor) — wastes work and may re-process

## Gotchas & Lessons Learned

- **`visited` is correctness, not optimization** — in a graph with cycles, skipping it causes infinite loops, not just inefficiency
- **Byte vs int literals strike again** — in `[][]byte` grids, compare against `'1'` (byte, value 49) NOT `1` (int). Go won't warn you; comparison is just always false. Same family of bug as Day 13 Decode String's `int(currChar)` vs `currChar - '0'`
- **Empty grid panic** — `len(grid[0])` on `[][]byte{}` is out-of-bounds. Always check outer length before indexing inner
- **Mark visited on enqueue, not on dequeue** — prevents duplicate enqueues. Learned in BFS on connected components
- **Clone Graph ordering**: check map → allocate → register → recurse. Any other order breaks cycles
- **Pointer keys vs value keys**: `map[*Node]*Node` is bulletproof for cloning; `map[int]*Node` (keyed on `node.Val`) only works if values are guaranteed unique
- **Undirected graphs need both directions in adjacency list** — easy to forget and only half-build the graph
- **Grid as graph** is a major perspective shift — no explicit `adj` needed. Coordinate math (`dr`, `dc`) replaces it

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 57 | Number of Connected Components | Component counting | Build adj list from edges, outer loop + BFS per unvisited node, count seeds | Mark visited on enqueue; isolated nodes still count |
| 58 | Number of Islands | Grid as implicit graph | Double-loop cells, BFS flood-fill on each unvisited `'1'`, count seeds | `'1'` byte (49) vs `1` int; empty-grid bounds guard |
| 59 | Clone Graph | DFS + memoization map | Map original→clone; check first, allocate, register, THEN recurse | Register clone in map BEFORE recursing — else cycles loop forever |
