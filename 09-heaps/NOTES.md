# Heaps / Priority Queues — Revision Notes

## How It Works
- A binary tree stored as a flat array where the parent is always "better" (smaller or larger) than its children
- Min-heap: parent <= children, smallest at root. Max-heap: parent >= children, largest at root
- No pointers — tree structure is implicit via index math: parent = `(i-1)/2`, left = `2*i+1`, right = `2*i+2`
- Insert: add at end, bubble UP. Extract: remove root, move last to top, bubble DOWN

```
Min-Heap array: [1, 3, 5, 7, 4]

       1          Index 0
      / \
     3   5        Index 1, 2
    / \
   7   4          Index 3, 4
```

## Complexity Cheat Sheet

| Operation | Time | Notes |
|-----------|------|-------|
| Insert (push) | O(log n) | Bubble up |
| Extract min/max (pop) | O(log n) | Bubble down |
| Peek min/max | O(1) | Read root |
| Build from array | O(n) | Bottom-up heapify |

## Key Patterns

### Top-K with Min-Heap of Size K
- **When to use:** Find K largest/most frequent elements without fully sorting
- **How it works:** Maintain a min-heap of size K. For each element, if it's bigger than the root (smallest in heap), swap it in. Root is always the Kth largest
- **Watch for:** Use min-heap for "K largest" (counterintuitive). The root is the gatekeeper — anything smaller gets rejected

### Simulation with Max-Heap
- **When to use:** Repeatedly process the largest/smallest elements (stone smashing, task scheduling)
- **How it works:** Insert all elements, then loop extracting the top 1-2 and reinserting results
- **Watch for:** Check heap size before extracting — may have 0 or 1 elements left

### Generic Heap with Priority/Value Separation
- **When to use:** When ordering key differs from stored data (e.g., sort by frequency, store element)
- **How it works:** `Entry[T]{Priority int, Value T}` — heap orders by Priority, returns Value
- **Watch for:** Use `PeekPriority()` to compare priorities, `Peek()` to get values — mixing them up causes wrong comparisons

## Gotchas & Lessons Learned
- Built heap from scratch rather than using `container/heap` — better for understanding internals
- Refactored from `MinHeap` struct → generic `Heap[T]` with compare function — reusable across all heap problems
- Added `(value, ok)` return pattern instead of magic `-1` for empty heap — safer and idiomatic Go
- `PeekPriority()` vs `Peek()` — need both when priority differs from value (Top K Frequent)
- Top K Frequent solved twice: Day 5 with sort O(n log n), Day 18 with heap O(n log k)

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 42 | Kth Largest Element | Min-heap size K | Push if size < K, else swap if bigger than root | Min-heap for K largest (not max-heap) |
| 43 | Last Stone Weight | Max-heap simulation | Extract two heaviest, reinsert difference if > 0 | Check size before final peek |
| 44 | Top K Frequent | Freq map + min-heap K | Count freq, heap of size K ordered by frequency | PeekPriority for frequency comparison |
