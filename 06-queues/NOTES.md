# Queues — Revision Notes

## How It Works
- FIFO (First In, First Out) — first element enqueued is first to dequeue
- Go: use `[]T` slice. Enqueue via `append`, dequeue via `q = q[1:]`
- Circular queue (ring buffer): fixed-size array with `front` and `rear` pointers using modular arithmetic
- BFS traversal uses a queue to process nodes level by level

```
Enqueue 1, 2, 3:
  Front → [1] [2] [3] ← Rear
Dequeue → returns 1
```

## Complexity Cheat Sheet

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Enqueue | O(1)* | O(1) | Amortized for slice-based |
| Dequeue | O(1) | O(1) | Reslice (but doesn't free memory) |
| Peek front | O(1) | O(1) | `queue[0]` |
| Peek rear | O(1) | O(1) | `queue[len(queue)-1]` |

**Circular queue:** All operations O(1) with no amortization — fixed memory.

## Key Patterns

### Two-Stack Queue
- **When to use:** Implement queue with stack-only operations (interview classic)
- **How it works:** inStack for push, outStack for pop. When outStack empty, pour all from inStack (reverse order = FIFO)
- **Watch for:** Clear inStack after pouring — forgetting this duplicates elements

### Monotonic Deque
- **When to use:** Sliding window maximum/minimum
- **How it works:** Deque stores indices in decreasing value order. Remove from back when new value is larger. Remove from front when index is out of window
- **Watch for:** Store **indices** not values (same lesson as monotonic stack). Compare values via `nums[deque[...]]`, remove indices when `i - front >= k`

### Circular Buffer (Ring Buffer)
- **When to use:** Fixed-size queue, OS schedulers, network buffers, producer-consumer
- **How it works:** Array with `front`, `rear`, `count`. Advance indices with `(idx + 1) % capacity`
- **Watch for:** Pointer semantics — be precise about "rear points to next write slot" vs "rear points to last written slot". Off-by-one from advancing rear before writing

### Frequency Map + Queue
- **When to use:** First unique element, streaming unique detection
- **How it works:** Queue holds candidates, frequency map tracks counts. Dequeue from front while count > 1
- **Watch for:** Check frequency at dequeue time, not enqueue time — counts may change after enqueue

## Gotchas & Lessons Learned
- Two-stack queue: must clear inStack after pouring to outStack — otherwise elements are duplicated
- Indices vs values confusion is a recurring theme: monotonic deque stores indices, compares via `nums[idx]`
- Circular buffer off-by-one: writing to `rear` then advancing vs advancing then writing changes everything
- `(rear - 1 + capacity) % capacity` to look back one step — the `+ capacity` handles the wrap when rear is 0

## Problem Summary

| # | Problem | Pattern | Approach (1-liner) | Tricky Part |
|---|---------|---------|-------------------|-------------|
| 21 | Queue Using Stacks | Two-stack pour | Push to inStack; pour to outStack on pop when empty | Clear inStack after pour |
| 22 | Sliding Window Maximum | Monotonic deque | Deque of indices in decreasing value order | Indices not values; remove when out of window |
| 23 | First Unique Character | Frequency + queue | Freq map + queue of candidates, drain non-uniques | Check freq at dequeue, not enqueue |
| 25 | Circular Queue | Ring buffer | Fixed array with front/rear/count, modular arithmetic | Rear semantics: write-then-advance vs advance-then-write |
