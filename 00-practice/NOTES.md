# Practice — Revision Notes

*Cross-topic problems from multiple topics. No single data structure — focus is on pattern recognition, cross-cutting lessons, and combining multiple data structures.*

## Cross-Cutting Lessons

- **Pattern recognition is the skill.** By Day 7, slow/fast pointer was identified independently without hints — proof that drilling patterns works
- **"Track min/max so far"** is a greedy technique that appears across many problems (stock prices, running statistics, etc.)
- **Frequency map** is the universal tool for intersection, union, and counting problems across arrays and strings
- **Recursive tree transformations** are often just one line — Invert Binary Tree reduces to swapping children and recursing. If you find yourself writing a lot of code for a tree problem, step back and think simpler
- **Multi-structure problems** (Merge K Lists) are about picking the right tool for each sub-task: heap for "give me the smallest of K things", dummy node for building a linked list
- **BFS level-order is a template** — Level Order Traversal, Right Side View, and many more all follow the same queue + level-size snapshot structure. The only thing that changes is what you collect per level
- **O(N log N) → O(N log K) optimization** — when K << N, keeping only K items in the heap instead of all N is a significant improvement. This pattern recurs in Top K problems and merge problems

## Gotchas & Lessons Learned
- Best Time to Buy and Sell Stock needed a walkthrough — the "track min so far" technique wasn't intuitive at first. Key: you don't need to find the min and max, just the max profit while scanning left to right
- Move Zeroes is slow/fast pointer, not a swap-based approach — recognizing the right pattern matters more than solving it
- Merge K Lists initial version inserted all N nodes into heap (O(N log N)). Optimized to insert only K heads and push next node after each extract (O(N log K))
- Merge K Lists nil pointer crash: must guard against nil lists in the input array before accessing `.Val`
- Invert Binary Tree first solved with helper function, then refactored to elegant single-expression: `root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)`

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 12 | Move Zeroes | Slow/fast pointer | Slow writes non-zeros, fast scans; fill rest with 0 | It's array dedup pattern, not swapping |
| 13 | Intersection of Two Arrays | Frequency map | Count freq of arr1, decrement for arr2 matches | Decrement to handle duplicates correctly |
| 14 | Best Time to Buy and Sell Stock | Running min + max profit | Track minPrice so far, update maxProfit = price - minPrice | Don't find global min/max; scan left-to-right |
| 46 | Merge K Sorted Lists | Min-heap + dummy node | Heap of K list heads, extract min, push its next | Nil lists in input; O(N log K) not O(N log N) |
