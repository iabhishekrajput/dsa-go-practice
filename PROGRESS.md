# DSA Learning Journey - Abhishek

**Started:** 2026-04-05
**Pace:** 1-2 hours/day
**Goal:** Master DSA from scratch
**Language:** Go (Golang)

---

## Curriculum Roadmap

### Phase 1: Foundations (Week 1-2)
- [x] Day 1: Big O Notation & Time/Space Complexity (skipped — already known)
- [x] Day 2: Arrays — the building block
- [x] Day 3: Arrays — sliding window & two pointers
- [x] Day 4: Strings — manipulation & pattern problems
- [x] Day 5: Hash Maps — your secret weapon (combined with Day 6)
- [x] Day 6: Hash Maps — frequency counting & grouping
- [x] Day 7: Practice day — mixed problems
- [x] Day 8: Linked Lists — why they exist (combined with Day 9)
- [x] Day 9: Linked Lists — reversal, cycle detection
- [x] Day 10: Stacks — LIFO magic
- [ ] Day 11: Queues — FIFO & variations
- [ ] Day 12: Stacks & Queues — real-world problems
- [ ] Day 13-14: Week 2 practice & review

### Phase 2: Trees & Recursion (Week 3-4)
- [ ] Recursion — thinking backwards
- [ ] Binary Trees — traversals (BFS, DFS)
- [ ] Binary Search Trees
- [ ] Tree problems — depth, path, LCA
- [ ] Heaps / Priority Queues
- [ ] Practice week

### Phase 3: Searching & Sorting (Week 5)
- [ ] Binary Search — beyond sorted arrays
- [ ] Sorting algorithms & when to use what
- [ ] Practice problems

### Phase 4: Graphs (Week 6-7)
- [ ] Graph representation & traversals
- [ ] BFS & DFS on graphs
- [ ] Shortest path algorithms
- [ ] Topological sort
- [ ] Practice problems

### Phase 5: Dynamic Programming (Week 8-9)
- [ ] DP intuition — memoization vs tabulation
- [ ] 1D DP problems
- [ ] 2D DP problems
- [ ] Classic DP patterns
- [ ] Practice problems

### Phase 6: Advanced (Week 10+)
- [ ] Tries
- [ ] Backtracking
- [ ] Greedy algorithms
- [ ] Bit manipulation
- [ ] Interview pattern revision

---

## Daily Log

### Day 1 — 2026-04-05 — Big O & Complexity
- **Status:** SKIPPED (already known)

### Day 2 — 2026-04-05 — Arrays: The Building Block
- **Status:** COMPLETED
- **Topics:** Array fundamentals, in-place operations, hash map lookup, sliding window
- **Problems solved:** 3/3
  1. Remove Duplicates from Sorted Array (slow/fast pointer)
  2. Two Sum (hash map)
  3. Max Sum Subarray of Size K (sliding window)
- **Notes:** Picked up all three patterns quickly. Clean idiomatic Go.

### Day 3 — 2026-04-05 — Arrays: Two Pointers & Sliding Window (Deeper)
- **Status:** COMPLETED
- **Topics:** Two-pointer convergence, variable-size sliding window
- **Problems solved:** 2/2
  4. Container With Most Water (two pointers from edges, move shorter side)
  5. Longest Substring Without Repeating Characters (variable sliding window + last-seen map)
- **Notes:** Handled the `abba` edge case (stale map entries). Needed one retry on #5 — initially used frequency counting instead of last-seen index approach.

### Day 4 — 2026-04-06 — Strings: Manipulation & Patterns
- **Status:** COMPLETED
- **Topics:** Two-pointer palindrome, string reversal, frequency counting
- **Problems solved:** 3/3
  6. Valid Palindrome (two pointers, skip non-alphanumeric)
  7. Reverse Words in a String (compact + reverse, reused slow/fast pointer)
  8. Valid Anagram (frequency counting with map)
- **Notes:** Palindrome solved first try. Reverse Words needed a fix on the compaction logic — recognized the slow/fast pointer pattern from Day 2. Anagram clean on first attempt.

### Day 5-6 — 2026-04-07 — Hash Maps: Grouping, Frequency & Sets
- **Status:** COMPLETED
- **Topics:** Grouping by computed key, frequency counting + sort, hash set for O(n) lookup
- **Problems solved:** 3/3
  9. Group Anagrams (sorted string as map key)
  10. Top K Frequent Elements (frequency map + sort unique keys)
  11. Longest Consecutive Sequence (hash set, only count from sequence starts)
- **Notes:** Common mistake on sort comparator (comparing indices vs values). Learned to work with unique keys from map, not raw input with duplicates. Used idiomatic map[int]struct{} for sets.

### Day 7 — 2026-04-07 — Practice Day: Mixed Problems
- **Status:** COMPLETED
- **Topics:** Slow/fast pointer, frequency map intersection, running min/max
- **Problems solved:** 3/3
  12. Move Zeroes (slow/fast pointer — identified pattern independently)
  13. Intersection of Two Arrays (frequency counting)
  14. Best Time to Buy and Sell Stock (running minimum + max profit)
- **Notes:** Correctly identified slow/fast for Move Zeroes without hints. Needed a walkthrough on Buy/Sell Stock — learned the "track min so far" technique.

### Day 8-9 — 2026-04-07 — Linked Lists: Fundamentals
- **Status:** COMPLETED
- **Topics:** Linked list reversal, cycle detection (Floyd's), merge sorted lists, dummy node pattern
- **Problems solved:** 3/3
  15. Reverse Linked List (prev/curr/next pointer flip)
  16. Linked List Cycle Detection (Floyd's tortoise & hare)
  17. Merge Two Sorted Lists (dummy node + zipper merge)
- **Notes:** Small bug returning curr instead of prev on reversal. Learned to check fast.Next != nil for safety in cycle detection. Used new nodes instead of rewiring in merge — learned to prefer rewiring for O(1) space.

### Day 10 — 2026-04-07 — Stacks: LIFO Magic
- **Status:** COMPLETED
- **Topics:** Stack basics, bracket matching, min tracking, monotonic stack
- **Problems solved:** 3/3
  18. Valid Parentheses (stack + closing→opening map)
  19. Min Stack (parallel min stack tracking)
  20. Daily Temperatures (monotonic stack — "next greater element" pattern)
- **Notes:** Learned make([]T, n) vs make([]T, 0, n) gotcha. Min Stack solved first try. Daily Temperatures needed help — key insight: result[top]=i-top, not result[i]. Monotonic stack is a new pattern to internalize.
