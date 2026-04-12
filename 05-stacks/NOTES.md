# Stacks — Revision Notes

## How It Works
- LIFO (Last In, First Out) — last element pushed is first to pop
- Go: use `[]int` slice as stack. `append` to push, `s[:len(s)-1]` to pop
- The call stack in recursion is literally a stack — each function call is a frame

```
Push 1, Push 2, Push 3:
  Top → [3]
        [2]
        [1]
Pop → returns 3
```

## Complexity Cheat Sheet

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Push | O(1)* | O(1) | Amortized (slice growth) |
| Pop | O(1) | O(1) | Reslice |
| Peek | O(1) | O(1) | `stack[len(stack)-1]` |
| IsEmpty | O(1) | O(1) | `len(stack) == 0` |

## Key Patterns

### Bracket Matching
- **When to use:** Validate nested structures (parentheses, HTML tags, etc.)
- **How it works:** Push opening brackets. On closing bracket, pop and check match. Stack must be empty at end
- **Watch for:** Use a closing→opening map for clean comparisons. Check stack empty before popping

### Monotonic Stack
- **When to use:** "Next greater/smaller element" problems, spans, temperatures
- **How it works:** Maintain stack in increasing/decreasing order. Pop elements that violate the order — popped elements found their answer
- **Watch for:** Store **indices** on the stack, not values. Result is `i - top`, not `result[i]`. This was a repeated lesson across Daily Temperatures, Sliding Window Max, and Next Greater Element

### Min Stack (Parallel Tracking)
- **When to use:** O(1) min/max retrieval alongside normal stack ops
- **How it works:** Maintain a shadow stack that tracks the current min at each level
- **Watch for:** Push to min stack even when equal (not just strictly less)

### Stack as Accumulator (RPN)
- **When to use:** Expression evaluation, calculator problems
- **How it works:** Numbers push to stack. Operators pop two, compute, push result
- **Watch for:** Operand order matters! First pop = right operand, second pop = left operand. `10 3 -` = 10-3, not 3-10

### Nested Structure Decoding
- **When to use:** Problems with nested brackets, nested repetition like `3[a2[c]]`
- **How it works:** On `[`, push current state (string + count) to stack and reset. On `]`, pop and combine
- **Watch for:** `int(rune)` gives Unicode code point (e.g., `'3'` → 51). Use `rune - '0'` for digit conversion. Multi-digit numbers: `num = num*10 + digit`

### Stack Simulation
- **When to use:** Process elements with conditional survival (collisions, cancellations)
- **How it works:** Push elements, pop when they conflict with incoming element. Track whether incoming element survives
- **Watch for:** Same-direction elements never collide — guard the inner loop. Track survival with a boolean

## Gotchas & Lessons Learned
- `make([]T, n)` creates n zero-valued elements; `make([]T, 0, n)` creates empty slice with capacity n — very different!
- Monotonic stack stores indices, not values — this was confusing on Daily Temperatures and Sliding Window Max
- RPN operand order: first pop is right, second is left. Got bitten by swapped subtraction/division
- `err != nil` means Atoi failed (token is operator), `err == nil` means success (token is number) — inverted this in RPN
- Asteroid collision: three separate conditions in the inner loop (smaller → destroy, equal → both destroy, same direction → no collision)

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 18 | Valid Parentheses | Bracket matching | Push opening, pop on closing, compare via map | Stack must be empty at the end |
| 19 | Min Stack | Parallel stack | Shadow min stack updated on every push/pop | Push min even when equal |
| 20 | Daily Temperatures | Monotonic stack | Stack of indices; pop when current temp > top | `result[top] = i - top`, not `result[i]` |
| 24 | Evaluate RPN | Stack accumulator | Push numbers; operators pop two, compute, push | First pop = right operand, second = left |
| 26 | Asteroid Collision | Stack simulation | Push right-movers; left-movers challenge stack top | Guard same-direction; track survival bool |
| 27 | Next Greater Element | Monotonic stack + map | Build next-greater map from nums2, lookup for nums1 | Combine monotonic stack with hash map |
| 28 | Decode String | Nested structure | Push (string, count) on `[`, pop and repeat on `]` | `rune - '0'` not `int(rune)` for digits |
