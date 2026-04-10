# Data Structures & Algorithms in Go

Comprehensive DSA implementations and problem solutions in Go, built from the ground up as a structured learning curriculum.

## Topics Covered

| # | Topic | Problems | Key Patterns |
|---|-------|----------|--------------|
| 01 | [Arrays](01-arrays/) | 5 | Two pointers, sliding window (fixed & variable), hash map lookup |
| 02 | [Strings](02-strings/) | 3 | Palindrome check, in-place reversal, frequency counting |
| 03 | [Hash Maps](03-hashmaps/) | 3 | Grouping by computed key, top-K frequency, hash set for O(n) lookup |
| 04 | [Practice](04-practice/) | 3 | Slow/fast pointer, running min/max, array intersection |
| 05 | [Linked Lists](05-linkedlists/) | 3 | Pointer reversal, Floyd's cycle detection, dummy node merge |
| 06 | [Stacks](06-stacks/) | 5 | Bracket matching, monotonic stack, min tracking, RPN evaluation, collision simulation |
| 07 | [Queues](07-queues/) | 4 | Two-stack queue, monotonic deque, circular buffer (ring buffer) |

**26 problems solved** across 7 topics.

## Project Structure

```
dsa/
├── 01-arrays/
│   ├── two_sum.go
│   ├── two_sum_test.go
│   └── ...
├── 02-strings/
├── 03-hashmaps/
├── 04-practice/
├── 05-linkedlists/
├── 06-stacks/
├── 07-queues/
├── PROGRESS.md          # Curriculum roadmap & daily log
└── README.md
```

Each problem consists of:
- **`problem.go`** — Solution implementation
- **`problem_test.go`** — Test cases covering edge cases, boundary conditions, and common pitfalls

## Running Tests

```bash
# Run all tests
go test ./... -v

# Run tests for a specific topic
go test ./06-stacks/ -v

# Run a single test
go test ./06-stacks/ -v -run TestEvalRPN
```

## Curriculum

The learning path follows a structured progression:

1. **Foundations** — Arrays, Strings, Hash Maps, Linked Lists, Stacks, Queues
2. **Trees & Recursion** — Binary Trees, BSTs, Heaps, DFS/BFS traversals
3. **Searching & Sorting** — Binary search, sorting algorithms
4. **Graphs** — Traversals, shortest path, topological sort
5. **Dynamic Programming** — Memoization, tabulation, classic DP patterns
6. **Advanced** — Tries, backtracking, greedy algorithms, bit manipulation

Full roadmap and daily progress in [PROGRESS.md](PROGRESS.md).

## Tech

- **Language:** Go 1.21+
- **Testing:** Standard `testing` package
- **Dependencies:** None — stdlib only
