package main 

import (
    "fmt"
    a "github.com/alegrey91/go-breaking-circular-dependencies/pkg/a"
    b "github.com/alegrey91/go-breaking-circular-dependencies/pkg/b"
)

/*
This project aims to explain how to avoid circular dependencies with golang
References:
- https://medium.com/@ishagirdhar/import-cycles-in-golang-b467f9f0c5a0
- Software Engineering with Golang
  (Chapter 2, pg 54, sec "Breaking circular dependencies via implicit interfaces")
*/
func main() {
    fmt.Println("Breaking Circular Dependencies")

    // How to use method from pkg b
    y := b.NewB()
    y.DoSomethingWithB()

    // How to use methods from pkg a passing through pkg b
    x := a.NewA()
    b.InvokeSomethingFromA(x)
}
