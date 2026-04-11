# Practice Day — Revision Notes

*This directory has mixed problems from multiple topics. No single data structure — focus is on pattern recognition and cross-cutting lessons.*

## Cross-Cutting Lessons

- **Pattern recognition is the skill.** By Day 7, slow/fast pointer was identified independently without hints — proof that drilling patterns works
- **"Track min/max so far"** is a greedy technique that appears across many problems (stock prices, running statistics, etc.)
- **Frequency map** is the universal tool for intersection, union, and counting problems across arrays and strings

## Gotchas & Lessons Learned
- Best Time to Buy and Sell Stock needed a walkthrough — the "track min so far" technique wasn't intuitive at first. Key: you don't need to find the min and max, just the max profit while scanning left to right
- Move Zeroes is slow/fast pointer, not a swap-based approach — recognizing the right pattern matters more than solving it

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 12 | Move Zeroes | Slow/fast pointer | Slow writes non-zeros, fast scans; fill rest with 0 | It's array dedup pattern, not swapping |
| 13 | Intersection of Two Arrays | Frequency map | Count freq of arr1, decrement for arr2 matches | Decrement to handle duplicates correctly |
| 14 | Best Time to Buy and Sell Stock | Running min + max profit | Track minPrice so far, update maxProfit = price - minPrice | Don't find global min/max; scan left-to-right |
