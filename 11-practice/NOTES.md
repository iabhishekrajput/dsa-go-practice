# Phase 2 Practice — Revision Notes

*Cross-topic practice problems combining trees, linked lists, and heaps. Focus is on applying learned patterns to new problems and combining multiple data structures.*

## Cross-Cutting Lessons

- **Recursive tree transformations** are often just one line — Invert Binary Tree reduces to swapping children and recursing. If you find yourself writing a lot of code for a tree problem, step back and think simpler
- **Multi-structure problems** (Merge K Lists) are about picking the right tool for each sub-task: heap for "give me the smallest of K things", dummy node for building a linked list
- **BFS level-order is a template** — Level Order Traversal, Right Side View, and many more all follow the same queue + level-size snapshot structure. The only thing that changes is what you collect per level
- **O(N log N) → O(N log K) optimization** — when K << N, keeping only K items in the heap instead of all N is a significant improvement. This pattern recurs in Top K problems and merge problems

## Gotchas & Lessons Learned
- Merge K Lists initial version inserted all N nodes into heap (O(N log N)). Optimized to insert only K heads and push next node after each extract (O(N log K))
- Merge K Lists nil pointer crash: must guard against nil lists in the input array before accessing `.Val`
- Invert Binary Tree first solved with helper function, then refactored to elegant single-expression: `root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)`

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 45 | Invert Binary Tree | Preorder recursion | Swap left/right, recurse into both children | Can be one line — don't overthink it |
| 46 | Merge K Sorted Lists | Min-heap + dummy node | Heap of K list heads, extract min, push its next | Nil lists in input; O(N log K) not O(N log N) |
| 47 | Binary Tree Right Side View | BFS level-order | Queue with level-size snapshot, take last per level | Left-deeper trees — rightmost isn't always a right child |
