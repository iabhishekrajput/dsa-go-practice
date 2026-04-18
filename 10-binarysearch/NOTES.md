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
| Rotated sorted search | O(log n) | O(1) | Identify sorted half, check if target is in it |
| Peak element | O(log n) | O(1) | Follow the upslope — a peak must exist on that side |
| Quickselect (kth element) | O(n) avg, O(n²) worst | O(log n) | Partition + recurse into only one side |

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

### Modified Binary Search (Rotated Array)
- **When to use:** Sorted array rotated at an unknown pivot; search for a target
- **How it works:** At each step, one half around `mid` is guaranteed sorted. Check `nums[left] <= nums[mid]` — if true, left half is sorted; otherwise right half is. Then check if target falls within the sorted half's range to decide which side to recurse into
- **Watch for:** Use `<=` (not `<`) for the sorted-half check — with a single element, `nums[left] == nums[mid]`. The comparison against `target` in the sorted half uses half-open bounds (e.g. `nums[left] <= target < nums[mid]`) so you don't double-count the mid

### Binary Search Without Sortedness (Peak Finding)
- **When to use:** Unsorted array, but some local property (peak, valley) lets you halve the search
- **How it works:** At `mid`, check slope: `nums[mid] < nums[mid+1]` means uphill, so a peak must exist to the right. Otherwise to the left (including `mid` itself)
- **Watch for:** Use `left < right` with `right = mid` (not `mid-1`) — `mid` might be the peak. Safe to access `nums[mid+1]` because `mid < right` when `left < right`

### Quickselect (Partition-Based Selection)
- **When to use:** Find kth smallest/largest without fully sorting
- **How it works:** Partition like quick sort. Pivot lands at final sorted position. If that matches target index, return; else recurse into only the side containing the target
- **Watch for:** kth largest maps to 0-indexed position `n-k`. Only recurse into one side — the other side is irrelevant. Average O(n) but worst O(n²) without random pivot

## Gotchas & Lessons Learned
- `mid = (left + right) / 2` can overflow in 32-bit languages. Always use `left + (right-left)/2`
- Half-open interval `[left, right)` with `left < right` is the cleanest template — `right` starts at `len(nums)`, and the answer is at `left` (or `right`, they're equal) when the loop ends
- Boundary search and insertion point are the same algorithm — "find the first element >= target"
- All three classic problems use the identical skeleton: only the comparison logic changes
- Rotated array: `nums[left] <= nums[mid]` must be `<=` not `<` — when `left == mid`, the "left half" is trivially sorted
- Peak finding: binary search doesn't require sortedness — it requires that some *invariant* lets you discard half. "A peak exists on the uphill side" is such an invariant
- Quickselect: the kth largest is at sorted-index `n-k`. When `pivotIdx == target`, early-return; otherwise recurse into the one side containing `target`
- `rand.Int() % (high-low)` excludes `high`. Use `rand.Intn(high-low+1) + low` for a uniform pick from `[low, high]`

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 48 | Binary Search | Exact match | Compare mid, go left/right, return on match | `left + (right-left)/2` for overflow safety |
| 49 | First Bad Version | Boundary search | No early return; bad → right=mid, good → left=mid+1 | Don't stop on first true — keep narrowing left |
| 50 | Search Insert Position | Insertion point | First element >= target; same as boundary search | Returns len(nums) if target > all elements |
| 54 | Search in Rotated Sorted Array | Modified binary search | Identify sorted half via `nums[left] <= nums[mid]`, check if target falls in that half | Must use `<=` (not `<`) for sorted-half check |
| 55 | Find Peak Element | Binary search on slope | Uphill (`nums[mid] < nums[mid+1]`) → go right; else `right = mid` | Not sorted — invariant is "a peak exists on uphill side" |
| 56 | Kth Largest Element | Quickselect | Partition, recurse only into side containing target index `n-k` | Map "kth largest" to 0-indexed `n-k`; only one-sided recursion |
