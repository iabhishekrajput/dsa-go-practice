# Linked Lists — Revision Notes

## How It Works
- Chain of nodes, each pointing to the next: `Val int` + `Next *ListNode`
- No random access — must walk from head to reach node N → O(n)
- Insertions/deletions at known positions are O(1) (just rewire pointers)
- Go: nodes are heap-allocated, garbage collected when unreachable

```
Head → [1|→] → [2|→] → [3|→] → nil
```

## Complexity Cheat Sheet

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Access by index | O(n) | O(1) | Walk from head |
| Insert at head | O(1) | O(1) | New node points to old head |
| Insert at tail | O(n) | O(1) | Walk to end first (O(1) with tail pointer) |
| Delete node | O(n) | O(1) | Find prev, rewire |
| Search | O(n) | O(1) | Linear scan |

## Key Patterns

### Pointer Reversal (prev/curr/next)
- **When to use:** Reverse a linked list or portion of it
- **How it works:** Track three pointers. Each step: save next, point curr back to prev, advance both
- **Watch for:** Return `prev` (not `curr`) — curr is nil when the loop ends

### Floyd's Cycle Detection (Tortoise & Hare)
- **When to use:** Detect if a linked list has a cycle
- **How it works:** Slow moves 1 step, fast moves 2 steps. If they meet, there's a cycle
- **Watch for:** Check `fast != nil && fast.Next != nil` before advancing fast — prevents nil dereference

### Dummy Node + Zipper Merge
- **When to use:** Merge two sorted lists, or any operation where the head might change
- **How it works:** Create a dummy node, build result by attaching nodes to dummy's tail. Return `dummy.Next`
- **Watch for:** Rewire existing nodes instead of creating new ones — O(1) space vs O(n)

## Gotchas & Lessons Learned
- Reversal: returning `curr` instead of `prev` is the #1 beginner bug. At loop end, `curr == nil` and `prev` is the new head
- Floyd's: always guard `fast.Next != nil` — skipping this causes nil pointer panic on even-length lists
- Merge: initially used `&ListNode{Val: ...}` to create new nodes. Better to rewire existing nodes with `curr.Next = smallerNode` for O(1) space

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 15 | Reverse Linked List | prev/curr/next flip | Save next, point curr→prev, advance both | Return prev, not curr |
| 16 | Linked List Cycle Detection | Floyd's tortoise & hare | Slow 1 step, fast 2 steps, check if they meet | Guard fast.Next != nil |
| 17 | Merge Two Sorted Lists | Dummy node + zipper | Dummy head, compare and attach smaller node | Rewire existing nodes, don't create new |
