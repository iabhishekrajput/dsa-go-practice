# Binary Search — Revision Notes

## How It Works
- Eliminate half the search space each step by comparing with the midpoint
- Requires a **monotonic property** — some condition that's false on one side and true on the other
- Classic setup: sorted array, but the concept applies anywhere you can halve the search space
- Half-open interval `[left, right)`: left inclusive, right exclusive. Loop: `left < right`. Converges when `left == right`

```
sorted: [1, 3, 5, 7, 9, 11, 15]    target = 9

lo=0        mid=3 (7<9)       hi=6
            lo=4   mid=5 (11>9)  hi=6
            lo=4   mid=4 (9==9) → found at index 4
```

## Complexity Cheat Sheet

| Variant | Time | Space | Notes |
|---------|------|-------|-------|
| Find exact match | O(log n) | O(1) | Early return on match |
| Find boundary | O(log n) | O(1) | No early return — keep narrowing |
| Search on answer space | O(log(range) * check) | O(1) | Binary search over possible answers |

## Key Patterns

### Exact Match
- **When to use:** Find a specific value in a sorted array
- **How it works:** Compare mid with target. Equal → return. Less → go right. Greater → go left
- **Watch for:** Use `mid = left + (right-left)/2` to avoid integer overflow

### Boundary Search (First True)
- **When to use:** Find first element satisfying a condition (first bad version, first >= target)
- **How it works:** If condition true → `right = mid` (answer might be here or earlier). If false → `left = mid + 1` (answer is strictly later). No early return — converge to the boundary
- **Watch for:** No `== target` check. The loop always runs to completion

### Insertion Point
- **When to use:** Find where a value belongs in sorted order
- **How it works:** Same as boundary search — find first element >= target. That index is the insertion point
- **Watch for:** If target exists, returns its index. If not, returns where it would go. Can return `len(nums)` for insert at end

## Gotchas & Lessons Learned
- `mid = (left + right) / 2` can overflow in 32-bit languages. Always use `left + (right-left)/2`
- Half-open interval `[left, right)` with `left < right` is the cleanest template — `right` starts at `len(nums)`, and the answer is at `left` (or `right`, they're equal) when the loop ends
- Boundary search and insertion point are the same algorithm — "find the first element >= target"
- All three problems use the identical skeleton: only the comparison logic changes

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 48 | Binary Search | Exact match | Compare mid, go left/right, return on match | `left + (right-left)/2` for overflow safety |
| 49 | First Bad Version | Boundary search | No early return; bad → right=mid, good → left=mid+1 | Don't stop on first true — keep narrowing left |
| 50 | Search Insert Position | Insertion point | First element >= target; same as boundary search | Returns len(nums) if target > all elements |
