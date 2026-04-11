# Arrays — Revision Notes

## How It Works
- Contiguous block of memory, indexed from 0
- Go slices: dynamic arrays backed by a fixed-size array, grow via `append` (amortized O(1))
- In-place operations avoid extra allocation by reusing the same backing array
- Subslicing (`nums[i:j]`) shares underlying memory — no copy

## Complexity Cheat Sheet

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Access by index | O(1) | O(1) | Direct memory offset |
| Append | O(1)* | O(1) | Amortized; O(n) when capacity doubles |
| Insert at index | O(n) | O(1) | Shift elements right |
| Delete at index | O(n) | O(1) | Shift elements left |
| Search (unsorted) | O(n) | O(1) | Linear scan |
| Search (sorted) | O(log n) | O(1) | Binary search |

## Key Patterns

### Slow/Fast Pointer (In-Place Writes)
- **When to use:** Remove/deduplicate elements in-place without extra space
- **How it works:** `slow` tracks the write position, `fast` scans ahead. Only write when `fast` finds a valid element
- **Watch for:** Off-by-one on return value (length vs index)

### Two Pointers (Convergence)
- **When to use:** Pairs in sorted data, or optimizing over two ends (e.g., max area)
- **How it works:** Start `left=0`, `right=n-1`, move inward based on a condition
- **Watch for:** Which pointer to move — always move the one that limits the current result (e.g., shorter side in container problem)

### Fixed Sliding Window
- **When to use:** Max/min/sum of all subarrays of size K
- **How it works:** Initialize window of size K, then slide: add right, subtract left
- **Watch for:** Initialize the first window separately before the slide loop

### Variable Sliding Window
- **When to use:** Longest/shortest subarray meeting a condition
- **How it works:** Expand `right`, shrink `left` when condition breaks. Track best result
- **Watch for:** Stale map entries — `abba` case. Use last-seen index, not just frequency. Update `left = max(left, lastSeen[char]+1)`

### Three-Reversal Rotation
- **When to use:** Rotate array by K positions in-place with O(1) space
- **How it works:** Reverse all → reverse first K → reverse rest. Swaps two halves without temp array
- **Watch for:** Normalize `k = k % len(nums)` first

## Gotchas & Lessons Learned
- `abba` edge case in variable sliding window: stale last-seen entries can shrink the window incorrectly
- Three-reversal trick is the same pattern used in Reverse Words (strings) — reversal is a universal building block
- Two-pointer convergence: always move the pointer that is currently the bottleneck

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 1 | Remove Duplicates from Sorted Array | Slow/fast pointer | Slow writes unique values, fast scans | Return slow+1 (length, not index) |
| 2 | Two Sum | Hash map lookup | Store num→index, check if complement exists | One pass is enough |
| 3 | Max Sum Subarray of Size K | Fixed sliding window | Init K-window, slide: add right, subtract left | Init first window before loop |
| 4 | Container With Most Water | Two pointers | Start at edges, move shorter side inward | Moving taller side can never improve area |
| 5 | Longest Substring Without Repeating | Variable sliding window | Expand right, shrink left via last-seen map | `abba` — update left with max, not just lastSeen |
| 29 | Rotate Array | Three-reversal | Reverse all, reverse [0,k-1], reverse [k,n-1] | Normalize k = k % n first |
