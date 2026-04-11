# Strings — Revision Notes

## How It Works
- Go strings are immutable byte slices; use `[]byte` or `[]rune` for in-place mutation
- `rune` = Unicode code point (int32); `byte` = single byte (uint8)
- String concatenation with `+` creates a new string each time — O(n) per concat
- `strings.Builder` for efficient multi-append string building

## Complexity Cheat Sheet

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Index access | O(1) | O(1) | By byte; rune access needs conversion |
| Concatenation | O(n+m) | O(n+m) | Creates new string |
| Substring | O(1) | O(1) | Shares backing array (like slices) |
| Convert to []byte | O(n) | O(n) | Full copy |
| len() | O(1) | O(1) | Byte length, not rune count |

## Key Patterns

### Two-Pointer Palindrome
- **When to use:** Check if string reads same forwards and backwards
- **How it works:** `left=0`, `right=n-1`, compare and converge. Skip non-alphanumeric characters
- **Watch for:** Case-insensitive comparison — normalize with `unicode.ToLower`

### In-Place Reversal (Compact + Reverse)
- **When to use:** Reverse words, rearrange parts of a string
- **How it works:** First compact (remove extra spaces), then reverse the whole thing, then reverse individual parts
- **Watch for:** The compaction step is a slow/fast pointer problem — same pattern as array dedup

### Frequency Counting
- **When to use:** Anagrams, character frequency comparison, finding duplicates
- **How it works:** Build `map[rune]int` (or `[26]int` for lowercase only), compare counts
- **Watch for:** Two-map vs single-map approaches — single map (increment/decrement) is cleaner

## Gotchas & Lessons Learned
- Reverse Words compaction logic is the slow/fast pointer pattern from Day 2 arrays — pattern recognition pays off
- `int(rune)` gives the Unicode code point (e.g., `'3'` → 51), not the digit value. Use `rune - '0'` for digit conversion
- Strings are immutable in Go — convert to `[]byte` for in-place operations

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 6 | Valid Palindrome | Two-pointer | Converge from edges, skip non-alphanumeric | Case normalization |
| 7 | Reverse Words in a String | Compact + reverse | Slow/fast compact, reverse all, reverse each word | Compaction is its own sub-problem |
| 8 | Valid Anagram | Frequency counting | Count chars in both strings, compare maps | Handle different lengths early |
