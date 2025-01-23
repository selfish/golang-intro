# Golang Basic Intro

- **Code first**
- **Move fast but keep it interactive**
- **Highlights Go's design philosophies** (simplicity, readability, concurrency) throughout

---

## 1. Quick Intro

5 minutes

- **Why Go?**
    - Simple syntax
    - Fast compile times
    - Concurrency built-in
    - Strong standard library
    - Used widely at scale (Docker, Kubernetes, etc.).
- **High-level differences from other popular languages** (e.g. Java, Python, C++).

---

## 2. Setting Up and Tooling

5 minutes

- **Go installation**: Mention `go mod` and module-based workflow introduced in Go 1.11+.
- **Tooling**:
    - `go run`, `go build`, `go test`

```shell
go clean -testcache
go test .
```

    - Basic `go mod init` usage

```shell
go mod init <packagename>
```

- **IDE/Editor tips**: Go extension in VSCode, or GoLand, etc., plus built-in tools like `gofmt`, `golint`.

```shell
gofmt <file>
gofmt -s <file>
gofmt -s -w <file>
gofmt -s -w .
```

---

## 3. Live Coding: Basic Syntax and Program Structure

10 minutes

- **Hello World**:
    - Show `package main`, `import "fmt"`, and the `main()` function.
    - Emphasize that `main()` is the entry point for executables.

```go
package main

import "fmt"

// Basic Hello World example showcasing package main, import, and main() function.
func main() {
	fmt.Println("Hello, Go World!")
}
```

- **Variables & Types**:
    - Basic declaration `var x int = 10` vs. short declaration `x := 10`.
    - Zero values (no uninitialized nonsense).
- **Functions**:
    - Return values (multiple returns), naming conventions.
    - Show how Go handles errors with multiple return values rather than exceptions.
- **Control flow**: `if`, `for` loops, `switch` statements. Mention lack of `while`.

---

## 4. Data Structures

10 minutes

- **Arrays vs. Slices**:
    - Slices are dynamic and widely used.
    - Quickly highlight make, append, and how slices relate to arrays under the hood.
- **Maps**:
    - Syntax, creation, iteration, checking for existence with `value, ok` pattern.
- **Structs**:
    - Demonstrate a simple struct and how to instantiate it.
    - Show how methods can be associated with types.
- **Pointers**:
    - Basic pointer usage, no pointer arithmetic.
    - Show how Go handles pass-by-value vs. reference to a pointer.

---

## 5. Concurrency & Goroutines

10 minutes

- **Goroutines**:
    - Show how easy it is to spin up concurrency with `go functionCall()`.
- **Channels**:
    - Basic example of sending/receiving from channels.
    - Emphasize channel usage for safe data sharing.
- **Quick mention**:
    - Select statement for multiplexing channels.
    - Sync package for WaitGroups if there's time.

---

## 6. Packages & Project Structure

5 minutes

- **Modular Code**:
    - Creating multiple packages, how imports work (`import “your_project/package”`).
    - `go mod` for dependencies.
- **Testing**:
    - Example of a small test using the standard `testing` package.
    - Running tests with `go test`.

---

## 7. Helpful Standard Library & Ecosystem

5 minutes

- **HTTP & Networking**:
    - Quick mention of `net/http` for servers.
    - Basic HTTP server in a few lines of code if time allows.
- **CLI**:
    - Mention `flag` package, or 3rd-party packages like "cobra" if relevant.
- **Community packages**:
    - The go.dev site, popular libraries (like gin-gonic for web frameworks).

---

## 8. Q&A / Wrap-up

10 minutes

- **Q&A**
- **Key resources**:
    - Official Go docs & tour (tour.golang.org)
    - Effective Go (official document)
    - Go by Example (gobyexample.com)

