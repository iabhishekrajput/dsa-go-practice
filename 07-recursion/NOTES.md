# Recursion — Revision Notes

## How It Works
- A function that calls itself to solve a smaller version of the same problem
- Every recursive solution has: **base case** (stop condition) + **recursive case** (delegate smaller problem)
- Each call creates a stack frame on the call stack — infinite recursion = stack overflow
- Think "delegation, not iteration" — trust the recursive call returns the right answer

```
factorial(4)
├── factorial(3)         stack: [f(4)]
│   ├── factorial(2)     stack: [f(4), f(3)]
│   │   ├── factorial(1) stack: [f(4), f(3), f(2)]
│   │   │   └── return 1           ← base case
│   │   └── return 2*1 = 2
│   └── return 3*2 = 6
└── return 4*6 = 24
```

## Complexity Cheat Sheet

| Pattern | Time | Space | Notes |
|---------|------|-------|-------|
| Single branch (linear) | O(n) | O(n) | One recursive call per level |
| Split in half | O(log n) | O(log n) | Fast exponentiation |
| Two branches | O(2^n) | O(n) | Without memoization |
| Divide & conquer | O(n log n) | O(n) | Merge sort |

Space is always **at least O(depth)** for the call stack.

## Key Patterns

### The Three Questions
For every recursive problem, answer:
1. **What's the base case?** (When do I stop?)
2. **What work do I do?** (What does this one call handle?)
3. **What do I delegate?** (What subproblem goes to the recursive call?)

### Fast Exponentiation
- **When to use:** Compute x^n efficiently
- **How it works:** Even n: `x^n = (x^(n/2))^2`. Odd n: `x^n = x * x^(n-1)`. Base: `x^0 = 1`
- **Watch for:** Store `halfPower` in a variable — calling `Power(x, n/2)` twice makes it O(n) instead of O(log n)

### Recursive Linked List Reversal
- **When to use:** Reverse a linked list without iteration
- **How it works:** Recurse to the end, then rewire: `head.Next.Next = head; head.Next = nil`
- **Watch for:** `newHead` (return value) is always the last node — it passes through unchanged. The rewiring happens through `head.Next.Next`, not through `newHead.Next`

### Divide and Conquer (Merge Sort)
- **When to use:** Sort, closest pair, many "split and combine" problems
- **How it works:** Split in half → recursively solve each half → merge results
- **Watch for:** The merge step is the real work — it's a two-pointer zipper merge (same as Merge Two Sorted Lists from linked lists)

## Gotchas & Lessons Learned
- Recursive reversal: `newHead` is NOT the tail of the reversed portion. It's always the last node of the original list, passed through unchanged. The node that should point back to you is `head.Next` (which is now the tail of the reversed sub-list)
- Fast exponentiation: computing `Power(x, n/2)` twice is a subtle O(n) bug. Store it once in a variable
- Don't trace every recursive call mentally — trust that the recursive call works (leap of faith)
- Negative exponents: handle once at the top with `1 / Power(x, -n)`, then recurse with positive n

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 30 | Power Function | Fast exponentiation | Even: square half, Odd: multiply by x, Base: n=0→1 | Store halfPower — don't call twice |
| 31 | Reverse List (Recursive) | Recursive reversal | Recurse to end; `head.Next.Next = head; head.Next = nil` | newHead vs tail confusion |
| 32 | Merge Sort | Divide & conquer | Split at mid, recurse both halves, two-pointer merge | Merge is the real work, not the split |
