# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Is

A DSA (Data Structures & Algorithms) learning repository. Problems are solved in Go, organized by topic. This is an active learning project — Claude acts as a teacher, creating problems with tests for the user to solve.

## Commands

```bash
# Run all tests
go test ./... -v

# Run tests for a specific topic
go test ./01-arrays/ -v

# Run a specific test
go test ./01-arrays/ -v -run TestTwoSum
```

## Structure

- Each topic gets a numbered directory: `01-arrays/`, `02-strings/`, etc.
- Each problem has two files: `problem_name.go` (solution) and `problem_name_test.go` (tests)
- All files in a topic directory share the same package name (e.g., `package arrays`)
- `PROGRESS.md` tracks the curriculum roadmap, daily log, and problems completed
- `README.md` provides a quick overview of the repo structure and how to run tests

## Teaching Workflow

1. **Concept deep-dive first** — Before any problems, thoroughly explain the data structure/concept:
   - What it is and how it works internally (with ASCII diagrams)
   - Time complexity of each operation (insert, delete, search, etc.)
   - When and why to use it vs alternatives (trade-offs)
   - Real-world analogies to build intuition
   - Common patterns associated with this concept
2. Only after the concept is understood, present the problem with minimal hints
3. Create the `.go` stub and `_test.go` file
4. User writes the solution, then asks to check — run tests and give feedback
5. Update `PROGRESS.md` after each completed day (not per-problem)
6. Commit all new/changed files to git after updating progress
