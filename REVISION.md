# Revision Notes Index

Quick-reference notes for interview prep, organized by topic. Each file has: how it works, complexity cheat sheet, key patterns, gotchas, and problem summaries.

## Topics

| # | Topic | Notes | Problems |
|---|-------|-------|----------|
| 01 | Arrays | [Notes](01-arrays/NOTES.md) | 6 |
| 02 | Strings | [Notes](02-strings/NOTES.md) | 3 |
| 03 | Hash Maps | [Notes](03-hashmaps/NOTES.md) | 3 |
| 04 | Practice | [Notes](04-practice/NOTES.md) | 3 |
| 05 | Linked Lists | [Notes](05-linkedlists/NOTES.md) | 3 |
| 06 | Stacks | [Notes](06-stacks/NOTES.md) | 7 |
| 07 | Queues | [Notes](07-queues/NOTES.md) | 4 |
| 08 | Recursion | [Notes](08-recursion/NOTES.md) | 3 |
| 09 | Trees & BSTs | [Notes](09-trees/NOTES.md) | 9 |
| 10 | Heaps / Priority Queues | [Notes](10-heaps/NOTES.md) | 3 |

**Total: 44 problems across 10 topics**

## Cross-Cutting Patterns

Patterns that appear across multiple topics — recognizing these is the real interview skill.

| Pattern | Where It Appears |
|---------|-----------------|
| Two pointers | Arrays (convergence, slow/fast), Strings (palindrome), Linked Lists (reversal, cycle) |
| Sliding window | Arrays (fixed & variable), Queues (monotonic deque) |
| Frequency map | Hash Maps, Strings (anagram), Arrays (two sum), Practice (intersection) |
| Monotonic stack/deque | Stacks (daily temps, next greater), Queues (sliding window max) |
| Dummy node | Linked Lists (merge), can apply to any head-may-change scenario |
| Divide and conquer | Recursion (merge sort), Trees (subtree problems) |
| Indices vs values | Stacks (monotonic), Queues (deque) — store indices, compare via array lookup |
| Postorder aggregation | Trees (depth, size), Recursion (children answer before parent) |
| Top-K with heap | Heaps (kth largest, top K frequent) — min-heap of size K as gatekeeper |
| Frequency map + heap | Hash Maps (count) + Heaps (select) — O(n log k) vs O(n log n) sort |
