# Sorting — Revision Notes

## How It Works

### Quick Sort
- Pick a **pivot**, partition array so smaller elements go left and larger go right, recurse on each side
- In-place (O(log n) stack space), average O(n log n), worst O(n²) with bad pivots
- **Random pivot** avoids worst case on sorted input — swap random element to low before partitioning

### Merge Sort (built on Day 14)
- Divide in half, sort each half, merge with two pointers
- Always O(n log n) but needs O(n) extra space. Stable sort

## Complexity Cheat Sheet

| Algorithm | Time (avg) | Time (worst) | Space | Stable? | Use when |
|-----------|-----------|-------------|-------|---------|----------|
| Quick Sort | O(n log n) | O(n²) | O(log n) | No | General purpose, in-place |
| Merge Sort | O(n log n) | O(n log n) | O(n) | Yes | Need guaranteed time, need stability |
| Heap Sort | O(n log n) | O(n log n) | O(1) | No | Need O(1) space with guaranteed time |
| Insertion Sort | O(n²) | O(n²) | O(1) | Yes | Small arrays, nearly sorted |

## Key Patterns

### Partition (Quick Sort Core)
- **When to use:** Divide array around a pivot; also used in quickselect, Dutch National Flag
- **How it works:** Two pointers from edges, swap elements on wrong side, place pivot at final position
- **Watch for:** Random pivot avoids O(n²) on sorted input. Swap pivot to `low` before partitioning — don't try to track its moving position

### Dutch National Flag (Three-Way Partition)
- **When to use:** Sort array with only 3 distinct values in one pass
- **How it works:** Three pointers — `low` (boundary of 0s), `mid` (scanner), `high` (boundary of 2s). Swap 0s left, 2s right, skip 1s
- **Watch for:** Don't advance `mid` after swapping with `high` — the swapped element hasn't been examined yet

### Sort Then Scan
- **When to use:** Merge intervals, find overlaps, remove duplicates, closest pair
- **How it works:** Sort first (O(n log n)), then single-pass scan (O(n)). Sorting makes adjacent elements comparable
- **Watch for:** After sorting intervals by start, overlapping intervals are always adjacent — compare current end with next start

## Gotchas & Lessons Learned
- Quick sort random pivot: must swap the random element to `low` position first, then partition normally. Swapping `nums[pivotIdx]` after partitioning is wrong because the pivot may have been moved by i/j swaps
- Merge intervals had 4 bugs in one attempt: (1) loop `i < len-1` drops last interval, (2) strict `>` should be `>=` for touching boundaries, (3) bounds check after index access, (4) overwrite `right` instead of `max(right, next)` for nested intervals
- Sort Colors: on case 0, both `low++` and `mid++` because anything at `low` is already processed. On case 2, only `high--` because the swapped element is unknown
- Writing custom sort for intervals (instead of `sort.Slice`) demonstrates understanding of internals — good interview signal

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 51 | Quick Sort | Partition + recurse | Random pivot to low, Hoare partition, recurse both sides | Swap pivot to low before partitioning, not after |
| 52 | Sort Colors | Dutch National Flag | Three pointers: 0→swap left, 1→skip, 2→swap right | Don't advance mid after swapping with high |
| 53 | Merge Intervals | Sort then scan | Sort by start, merge if current end >= next start, take max end | Touching boundaries (>=), nested intervals (max end), include last interval |
