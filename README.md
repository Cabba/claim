claim
===

Defensive programming in Go.

This library provides a simple and lightweight way to implement defensive programming techniques (preconditions and postconditions) in Go.

## Install

```bash
go get github.com/Cabba/claim
```

## Usage

### Preconditions

Use `claim.Pre()` at the beginning of a function to define preconditions.  These are conditions that *must* be true before the function executes. If a precondition fails, the program will panic (unless running in `unclaim` mode – see below).

```go
func sum(a, b int) int {
    claim.Pre(b != 0) 

    return a + b
}
```

### Postconditions

Use `claim.Post()` with `defer` to define postconditions. 
These are conditions that *must* be true after the function has executed.  
They verify the function's results. If a postcondition fails, the program will panic (unless running in `unclaim` mode).

```go
func discountPrice(price float32) float32 {
    discountedPrice := 0

    defer claim.Post(discountedPrice >= 0)

    // make your discount computations ...

    return discountedPrice
}

func absoluteValue(x int) int {
    result := x
    if x < 0 {
        result = -x
    }

    // you can use claim.Post just before returning
    claim.Post(result >= 0)
    return result
}
```

### State Change Postconditions

`claim.Post()` is also useful for verifying state changes:

```go
func (d *Door) open() {
    defer claim.Post(d.state == StateDoorOpen)
    
    // perform state mutation
}
```

### `unclaim` Mode

If you're lighthearted and dislike panicking use `unclaim` mode. 
In this mode, `claim` will print an error message and a stack trace to `claim.UnclaimWriter` (which defaults to `os.Stderr`). 
This allows you to continue execution and see where the violations occurred, without halting your program.

To enable `unclaim` mode, build your code with the `unclaim` build tag:

```bash
go build -tags unclaim .
```

You can customize the output writer by setting `claim.UnclaimWriter`:

```go
func main() {
    claim.UnclaimWriter = os.Stdout // Example: write to standard output

    // ... your code ...
}
```
