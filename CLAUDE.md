# CLAUDE.md

<!-- code-review-graph MCP tools -->
## MCP Tools: code-review-graph

**IMPORTANT: This project has a knowledge graph. ALWAYS use the
code-review-graph MCP tools BEFORE using Grep/Glob/Read to explore
the codebase.** The graph is faster, cheaper (fewer tokens), and gives
you structural context (callers, dependents, test coverage) that file
scanning cannot.

### When to use graph tools FIRST

- **Exploring code**: `semantic_search_nodes` or `query_graph` instead of Grep
- **Understanding impact**: `get_impact_radius` instead of manually tracing imports
- **Code review**: `detect_changes` + `get_review_context` instead of reading entire files
- **Finding relationships**: `query_graph` with callers_of/callees_of/imports_of/tests_for
- **Architecture questions**: `get_architecture_overview` + `list_communities`

Fall back to Grep/Glob/Read **only** when the graph doesn't cover what you need.

### Key Tools

| Tool | Use when |
|------|----------|
| `detect_changes` | Reviewing code changes — gives risk-scored analysis |
| `get_review_context` | Need source snippets for review — token-efficient |
| `get_impact_radius` | Understanding blast radius of a change |
| `get_affected_flows` | Finding which execution paths are impacted |
| `query_graph` | Tracing callers, callees, imports, tests, dependencies |
| `semantic_search_nodes` | Finding functions/classes by name or keyword |
| `get_architecture_overview` | Understanding high-level codebase structure |
| `refactor_tool` | Planning renames, finding dead code |

### Workflow

1. The graph auto-updates on file changes (via hooks).
2. Use `detect_changes` for code review.
3. Use `get_affected_flows` to understand impact.
4. Use `query_graph` pattern="tests_for" to check coverage.


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

## Teaching Workflow

1. Explain the concept with real-world analogies
2. Present the problem with hints
3. Create the `.go` stub and `_test.go` file
4. User writes the solution, then asks to check — run tests and give feedback
5. Update `PROGRESS.md` after each completed problem
