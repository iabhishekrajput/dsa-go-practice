# Hash Maps — Revision Notes

## How It Works
- Key-value store with O(1) average lookup via hash function
- Go `map[K]V`: built-in hash map, unordered, safe for concurrent reads but not writes
- Hash collisions handled internally (chaining in Go's implementation)
- `map[int]struct{}` is the idiomatic Go hash set (zero-size value)

## Complexity Cheat Sheet

| Operation | Time (avg) | Time (worst) | Space | Notes |
|-----------|-----------|-------------|-------|-------|
| Insert | O(1) | O(n) | O(1) | Worst case: all keys collide |
| Lookup | O(1) | O(n) | O(1) | Use `val, ok := m[key]` idiom |
| Delete | O(1) | O(n) | O(1) | `delete(m, key)` |
| Iterate | O(n) | O(n) | O(1) | Random order in Go |

## Key Patterns

### Grouping by Computed Key
- **When to use:** Group items that share a property (anagrams, same pattern, etc.)
- **How it works:** Compute a canonical key from each item (e.g., sorted string), use as map key to group
- **Watch for:** Key type must be comparable in Go — use `string`, not `[]byte` or `[]int` as map key

### Frequency Counting + Sort
- **When to use:** Top-K frequent elements, most/least common
- **How it works:** Build frequency map, extract unique keys, sort by frequency, take top K
- **Watch for:** Sort comparator should compare values (frequencies), not keys. Common mistake: comparing indices vs values

### Hash Set for O(n) Lookup
- **When to use:** Existence checks, deduplication, finding sequences/pairs
- **How it works:** Store elements in `map[int]struct{}`, check membership in O(1)
- **Watch for:** For consecutive sequences, only start counting from sequence starts (where `num-1` doesn't exist) — this is the optimization that makes it O(n)

## Gotchas & Lessons Learned
- Sort comparator bug: comparing indices vs values — always double-check what you're sorting by
- Extract unique keys from map first, then sort — don't sort raw input that may have duplicates
- `map[int]struct{}` is idiomatic Go for sets; `map[int]bool` works but wastes a byte per entry
- Longest Consecutive Sequence: the key insight is only counting from sequence starts — without this, it's O(n log n) or worse

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 9 | Group Anagrams | Grouping by key | Sort each string as map key, group values | Key must be string type (not []byte) |
| 10 | Top K Frequent Elements | Frequency + sort | Frequency map → unique keys → sort by freq → take K | Sort comparator: compare frequencies, not keys |
| 11 | Longest Consecutive Sequence | Hash set | Set lookup; only count from sequence starts | Start detection: skip if num-1 exists in set |
