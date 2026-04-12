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
- [x] Day 11: Queues — FIFO & variations
- [x] Day 12: Stacks & Queues — real-world problems
- [x] Day 13-14: Week 2 practice & review

### Phase 2: Trees & Recursion (Week 3-4)
- [x] Recursion — thinking backwards
- [x] Binary Trees — traversals (BFS, DFS)
- [x] Binary Search Trees
- [x] Tree problems — depth, path, LCA
- [x] Heaps / Priority Queues
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

### Day 11 — 2026-04-07 — Queues: FIFO & Variations
- **Status:** COMPLETED
- **Topics:** Queue concept deep-dive, queue using stacks, monotonic deque, frequency + queue combo
- **Problems solved:** 3/3
  21. Queue Using Stacks (two-stack pour pattern)
  22. Sliding Window Maximum (monotonic deque — hard)
  23. First Unique Character (frequency map + queue)
- **Notes:** Forgot to clear inStack after pouring in queue-using-stacks. Sliding Window Max needed help with indices-vs-values confusion — same lesson as Daily Temperatures. First Unique done cleanly first try.

### Day 12 — 2026-04-10 — Stacks & Queues: Real-World Problems
- **Status:** COMPLETED
- **Topics:** Expression evaluation, ring buffer design, stack simulation
- **Problems solved:** 3/3
  24. Evaluate Reverse Polish Notation (stack as accumulator — operand order gotcha)
  25. Circular Queue / Ring Buffer (modular arithmetic, pointer semantics)
  26. Asteroid Collision (stack simulation with chain reactions)
- **Notes:** RPN had two bugs: inverted error check and swapped operand order. Circular queue off-by-one from advancing rear before writing. Asteroid collision needed guard for same-direction asteroids and survival tracking for left-movers that clear the stack.

### Day 13 — 2026-04-10 — Week 2 Practice & Review
- **Status:** COMPLETED
- **Topics:** Monotonic stack + hash map combo, nested stack decoding, array reversal trick
- **Problems solved:** 3/3
  27. Next Greater Element (monotonic stack + map lookup — solved first try)
  28. Decode String (stack of string/count frames — rune-to-int bug with `int(currChar)` vs `currChar - '0'`)
  29. Rotate Array (three-reversal pattern — solved first try)
- **Notes:** Pattern recognition is strong — Next Greater Element and Rotate Array both solved cleanly on first attempt. Decode String had a subtle rune-to-digit conversion bug. Phase 1 Foundations complete!

### Day 14 — 2026-04-11 — Recursion: Thinking Backwards
- **Status:** COMPLETED
- **Topics:** Fast exponentiation, recursive linked list reversal, divide-and-conquer (merge sort)
- **Problems solved:** 3/3
  30. Power Function (fast exponentiation O(log n) — solved first try)
  31. Reverse Linked List Recursive (head.Next.Next rewire — initially attached to newHead instead of tail)
  32. Merge Sort (divide-and-conquer with two-pointer merge — solved first try)
- **Notes:** Power and Merge Sort solved cleanly on first attempt. Recursive reversal tripped on confusing newHead (always the last node) with the tail of the reversed portion — key insight is that head.Next still points to the right node after recursion.

### Day 15 — 2026-04-11 — Binary Trees: Traversals (BFS, DFS)
- **Status:** COMPLETED
- **Topics:** DFS vs BFS concept deep-dive, inorder traversal, level-order (BFS), recursive depth calculation
- **Problems solved:** 3/3
  33. Inorder Traversal (recursive Left→Root→Right with slice concat — solved first try)
  34. Level Order Traversal (BFS with queue, level-size snapshot trick — solved first try)
  35. Maximum Depth (postorder pattern: 1 + max(left, right) — solved first try)
- **Notes:** All three solved first try. DFS/BFS distinction clicked after side-by-side stack-vs-queue walkthrough. Used Go 1.22 range-over-int in BFS. Recursive tree thinking coming naturally from Day 14 foundation.

### Day 16 — 2026-04-11 — Binary Search Trees
- **Status:** COMPLETED
- **Topics:** BST property, range validation, BST search, inorder = sorted order
- **Problems solved:** 3/3
  36. Validate BST (pass min/max range down — solved first try)
  37. Search in BST (binary search on tree — solved first try)
  38. Kth Smallest Element (inorder with pointer counter, early termination — solved first try)
- **Notes:** All three first-try solves. Used optimal pointer-based counter for Kth Smallest instead of collecting full inorder. BST patterns feel natural — range narrowing for validation, directional search, inorder for sorted access.

### Day 17 — 2026-04-12 — Tree Problems: Depth, Path, LCA
- **Status:** COMPLETED
- **Topics:** Mirror comparison, carry-state-down recursion, postorder bubbling for LCA
- **Problems solved:** 3/3
  39. Symmetric Tree (parallel subtree mirror check — solved first try)
  40. Path Sum (subtract-as-you-go, check at leaf — solved first try)
  41. Lowest Common Ancestor (postorder: both sides non-nil = LCA — solved first try)
- **Notes:** All three first-try solves. Symmetric tree combined nil checks cleanly into one condition. Path Sum correctly identified leaf-only check. LCA pattern is elegant — the split detection via left/right non-nil is the key insight.

### Day 18 — 2026-04-12 — Heaps / Priority Queues
- **Status:** COMPLETED
- **Topics:** Heap internals, min/max heap, generic heap with Entry[T], top-K pattern, simulation with heap
- **Problems solved:** 3/3
  42. Kth Largest Element (min-heap of size K — solved first try)
  43. Last Stone Weight (max-heap simulation — solved first try)
  44. Top K Frequent Elements (frequency map + min-heap of size K — solved first try)
- **Notes:** All three first-try solves. Built heap from scratch instead of using container/heap — then refactored into generic `Heap[T]` with `Entry[T]{Priority, Value}` and `PeekPriority()`. Revisited Top K Frequent from Day 5 (sort) with heap approach (O(n log k) vs O(n log n)). 15-problem first-try streak across Days 15-18.
