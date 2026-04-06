# DSA Go Practice

A hands-on DSA (Data Structures & Algorithms) learning journey, solved in Go.

## Structure

```
01-arrays/      — Two pointers, sliding window, hash map lookup
02-strings/     — Palindromes, reversal, frequency counting
03-hashmaps/    — (coming soon)
...
```

Each problem has a solution file (`problem.go`) and tests (`problem_test.go`).

## Running Tests

```bash
# All tests
go test ./... -v

# Single topic
go test ./01-arrays/ -v

# Single problem
go test ./01-arrays/ -v -run TestTwoSum
```

## Progress

Tracked in [PROGRESS.md](PROGRESS.md).
