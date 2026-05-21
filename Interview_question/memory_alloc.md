# Go Closures, Escape Analysis, Stack vs Heap — Interview Notes

## 1. Problem Statement

Consider this Go code:

```go
package main

import "fmt"

const a = 10
var p = 100

func Outer() func() {
    money := 100
    age := 30

    fmt.Println("Age =", age)

    show := func() {
        money = money + p + a
        fmt.Println(money)
    }

    return show
}

func call() {
    incr1 := Outer()
    incr1()
    incr1()

    incr2 := Outer()
    incr2()
    incr2()
}
```

---

# 2. Important Concepts Involved

This example demonstrates:

- Closures
- Anonymous functions
- Escape Analysis
- Stack vs Heap memory
- Variable lifetime
- Garbage Collection
- Function scope

---

# 3. Compilation Phase vs Runtime Phase

---

## Compile Time

During compilation:

### Code Segment Stores

- Function definitions
- Constants
- Instructions

Examples:

```go
const a = 10
```

```go
func Outer() {}
```

```go
func call() {}
```

Anonymous function:

```go
show := func() {}
```

also becomes part of the executable code segment.

---

# Runtime Phase

During execution:

## Data Segment

Global variables are stored here.

Example:

```go
var p = 100
```

This exists throughout the entire program lifetime.

---

# 4. Stack Memory

Normally local variables are stored in the stack.

Example:

```go
func test() {
    x := 10
}
```

Here:

- `x` is placed in stack memory
- When `test()` finishes:
  - stack frame destroyed
  - `x` disappears

Stack memory is:

- fast
- automatically managed
- short-lived

---

# 5. What Is a Closure?

A closure is:

> A function that remembers and accesses variables from its outer scope even after the outer function has finished execution.

Example:

```go
func Outer() func() {
    money := 100

    show := func() {
        fmt.Println(money)
    }

    return show
}
```

The inner function `show` uses `money` from the outer function.

This is called:

# Variable Capture

---

# 6. Why Does `money` Move to the Heap?

Inside `Outer()`:

```go
money := 100
```

Normally this should stay in stack memory.

BUT:

```go
return show
```

The returned function still needs access to `money`.

If `money` stayed in the stack:

- stack frame would be destroyed after `Outer()` returns
- `money` would become invalid

Therefore Go compiler performs:

# Escape Analysis

The compiler detects:

> “This variable escapes the function scope.”

So Go allocates `money` in:

# Heap Memory

instead of stack memory.

---

# 7. Why `age` Does NOT Escape

```go
age := 30
fmt.Println(age)
```

`age` is:

- used only inside `Outer()`
- not returned
- not captured by closure

Therefore:

- it stays in stack memory
- destroyed after `Outer()` ends

---

# 8. Heap Allocation Flow

---

## First Call

```go
incr1 := Outer()
```

Creates:

```go
money = 100
```

Heap allocation happens.

Returned closure keeps reference to that heap variable.

---

## First Execution

```go
incr1()
```

Calculation:

```go
money = 100 + 100 + 10
      = 210
```

Output:

```text
210
```

---

## Second Execution

```go
incr1()
```

Same heap variable reused:

```go
money = 210 + 100 + 10
      = 320
```

Output:

```text
320
```

---

# 9. Why `incr2` Starts Again From 100

```go
incr2 := Outer()
```

This creates:

- a completely new closure
- a new captured environment
- another independent heap variable

So:

```text
210
320
210
320
```

---

# 10. Stack vs Heap Comparison

| Stack               | Heap                    |
| ------------------- | ----------------------- |
| Fast                | Slower                  |
| Automatically freed | Garbage Collector frees |
| Short lifetime      | Long lifetime           |
| Function-local      | Shared/persistent       |
| Small memory        | Large memory            |

---

# 11. What Is Escape Analysis?

Escape analysis is a compiler optimization technique.

The Go compiler decides:

- Should variable stay in stack?
- Or move to heap?

---

## Variables Escape When

- Returned from function
- Referenced by closure
- Referenced outside current scope
- Lifetime exceeds function execution

---

# 12. Garbage Collector’s Role

Heap memory is managed by:

# Garbage Collector (GC)

When no closure references `money` anymore:

- GC detects unused memory
- memory gets cleaned automatically

---

# 13. Interview Questions & Answers

---

## Q1. What is a closure in Go?

### Answer

A closure is a function that captures and remembers variables from its outer lexical scope even after the outer function has returned.

---

## Q2. Why does `money` escape to heap?

### Answer

Because the returned anonymous function still references `money` after `Outer()` returns. Therefore its lifetime exceeds the function scope, forcing heap allocation.

---

## Q3. What is escape analysis?

### Answer

Escape analysis is a compiler process used to determine whether variables should be allocated on the stack or heap depending on their lifetime and usage.

---

## Q4. Why does `age` stay in stack?

### Answer

Because `age` is only used inside `Outer()` and does not escape the function scope.

---

## Q5. Why is heap slower than stack?

### Answer

Heap memory requires dynamic allocation and garbage collection, while stack allocation is simple pointer movement and automatic cleanup.

---

# 14. Important Go Runtime Insight

This example demonstrates:

- closures create stateful functions
- Go automatically handles memory allocation
- compiler and GC work together
- heap allocation is often caused by closures

---

# 15. Key Takeaways

## Remember These Points

### Stack

- short-lived
- function-local
- very fast

### Heap

- long-lived
- GC-managed
- used when variables escape

### Closures

- capture outer variables
- preserve state between calls

### Escape Analysis

- compiler decides stack vs heap

---

# 16. Expected Output

```text
Age = 30
210
320
Age = 30
210
320
```

---

# 17. One-Line Interview Summary

> “`money` moves to the heap because the returned closure keeps referencing it after `Outer()` returns, causing the variable to escape its original stack frame.”
