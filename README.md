# Getting started with go

## First project

Run
```sh
    cd ./your-repos
    mkdir -p go-project && cd go-project
    go mod init example.com/your-pkg-repo
```
Open a `hello.go` file in editor and write code

```go
    package main

    import "fmt"
    import "rsc.io/quote"

    func main() {
        fmt.Println(quote.Go())
    }
```

Here, we are
- declaring a package `main`
- importing a package from go standard library- `fmt`
- importing a package from `pkg.go.dev` - `quote`
- calling code from _external package_ inside _**main function** of our package_

Run
```sh
go mod tidy
# to "tidy" up `go.mod` file. In the sense of syncing our modules, removed or newly added
go run .
# or go run ./hello.go
```
## Pointers

Pointers store memory location of an address in memory
Useful in passing data by reference rather than value[^1]

`&my_var` returns memory __address of__ my_var
`*T` is said to be a pointer to some type T which has a certain address in memory
This memory address is stored in the pointer
Dereferencing a *T pointer will yeild type T and is used to get the data stored at that memory location or to assign new data i.e. mutate the memory location
> Zero value for pointers is `nil` i.e. uninitialized pointers will store nil

__No pointer arithmetic__ in Go

---

[^1]: which ends up being copied, and copying is expensive for large data structures

