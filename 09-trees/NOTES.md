# Trees — Revision Notes

## How It Works
- Hierarchical structure: each node has a value + pointers to left and right children
- Binary tree: each node has **at most 2 children**
- Root = top node (no parent), Leaf = node with no children, Depth = edges from root to node
- Go: `type TreeNode struct { Val int; Left, Right *TreeNode }`

```
         1          ← root (depth 0)
        / \
       2   3        ← depth 1
      / \   \
     4   5   6      ← depth 2 (leaves)
```

## Complexity Cheat Sheet

| Operation | Balanced Tree | Skewed Tree | Notes |
|-----------|-------------|-------------|-------|
| DFS traversal | O(n) | O(n) | Visit every node |
| BFS traversal | O(n) | O(n) | Visit every node |
| DFS space | O(log n) | O(n) | Call stack = height |
| BFS space | O(w) | O(1) | Queue holds widest level |
| BST search | O(log n) | O(n) | Balanced vs degenerate |
| BST insert | O(log n) | O(n) | Same as search |

## Key Patterns

### DFS Traversals (Recursive)
- **When to use:** Most tree problems — depth, paths, subtree properties
- **How it works:** Three variants, differing only in when you process the node:
  - **Preorder:** Root → Left → Right (process before children)
  - **Inorder:** Left → Root → Right (sorted order for BST)
  - **Postorder:** Left → Right → Root (children answer before parent)
- **Watch for:** All three are the same code — only the position of `visit(node)` changes

### BFS / Level-Order Traversal
- **When to use:** Process level by level, shortest path in unweighted tree, level grouping
- **How it works:** Queue of nodes. Key trick: snapshot `levelSize = len(queue)` at start of each level, process exactly that many nodes
- **Watch for:** Without the level-size snapshot, you mix nodes from different levels since you enqueue children while processing

### Postorder for Aggregation
- **When to use:** Max depth, subtree size, tree diameter — anything where parent needs children's answers
- **How it works:** Recurse left, recurse right, combine: `1 + max(leftDepth, rightDepth)`
- **Watch for:** Base case is `nil → 0` (not 1). The "1 +" accounts for the current node

---

## BST-Specific Patterns

### BST Property
For every node: **all left subtree values < node < all right subtree values**. This gives O(log n) search on balanced trees.

### Range Validation (Validate BST)
- **When to use:** Check if a binary tree satisfies BST property
- **How it works:** Pass `(min, max)` range down. Go left → node becomes new max. Go right → node becomes new min
- **Watch for:** Checking only parent vs child is NOT enough — must validate against all ancestors. `5 → right:6 → left:3` is invalid (3 < 5) but passes parent-only check

### BST Search
- **When to use:** Find a value in a BST
- **How it works:** Compare with root: equal → found, less → go left, greater → go right. O(log n) on balanced tree
- **Watch for:** Returns the node (with its subtree), not just a boolean

### Inorder = Sorted Order
- **When to use:** Kth smallest, range queries, sorted output from BST
- **How it works:** Inorder traversal of BST visits nodes in ascending order
- **Watch for:** For kth smallest, use a pointer-based counter (`*int`) with early termination instead of collecting the full slice — stops as soon as kth node is visited

---

## Tree Problem Patterns

### Mirror/Symmetry Check
- **When to use:** Check if a tree is symmetric, or compare two subtrees
- **How it works:** Compare two nodes in parallel — outer children mirror outer, inner children mirror inner
- **Watch for:** Three base cases: both nil (true), one nil (false), values differ (false)

### Carry State Downward (Path Sum)
- **When to use:** Root-to-leaf path problems — sum, path collection, constraints
- **How it works:** Subtract/accumulate as you recurse down, check condition at leaf nodes only
- **Watch for:** Leaf = both children nil. A node with one child is NOT a leaf — don't check the sum there

### Lowest Common Ancestor (Postorder Bubbling)
- **When to use:** Find deepest shared ancestor of two nodes
- **How it works:** Recurse both subtrees. If both return non-nil, current node is the LCA. If only one side returns non-nil, bubble that up
- **Watch for:** A node can be its own ancestor (if p is ancestor of q, LCA is p)

## Gotchas & Lessons Learned
- DFS vs BFS: DFS uses stack (call stack or explicit), BFS uses queue. Not new algorithms — just stack vs queue applied to tree exploration
- Level-order grouping: must snapshot queue length before processing — without it, newly enqueued children mix into the current level
- BST validation trap: only checking parent vs child misses grandparent violations. Always pass the valid range
- `reflect.DeepEqual` treats `nil` and `[]int{}` as different — return `nil` for empty trees if tests expect it
- Pointer-based counter (`*int`) is idiomatic Go for sharing mutable state across recursive calls without collecting into a slice

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 33 | Inorder Traversal | DFS (recursive) | Left → append root → Right, concat slices | Base case returns `[]int{}` not `nil` |
| 34 | Level Order Traversal | BFS (queue) | Queue + level-size snapshot, group per level | Snapshot len(queue) before inner loop |
| 35 | Maximum Depth | Postorder aggregation | `1 + max(left, right)`, base nil→0 | Postorder: children answer before parent |
| 36 | Validate BST | Range narrowing | Pass (min, max), narrow on each step | Check all ancestors, not just parent |
| 37 | Search in BST | BST binary search | Equal→found, less→left, greater→right | Return subtree, not just boolean |
| 38 | Kth Smallest in BST | Inorder + counter | Inorder with `*int` counter, stop at k=0 | Pointer counter for early termination |
| 39 | Symmetric Tree | Mirror check | Compare two subtrees: outer↔outer, inner↔inner | Three base cases: both nil, one nil, values differ |
| 40 | Path Sum | Carry state down | Subtract as you go, check remaining==val at leaf | Leaf = both children nil, not just one |
| 41 | Lowest Common Ancestor | Postorder bubbling | Recurse both sides; both non-nil = LCA | Node can be its own ancestor |
