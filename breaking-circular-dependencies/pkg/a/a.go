package a

import (
    "fmt"
    "github.com/alegrey91/go-breaking-circular-dependencies/pkg/b"
)

type A struct {}

func NewA() *A {
    a := A{}
    return &a
}

func (a A) DoSomethingWithA() {
  fmt.Println("This is a")
}

func InvokeSomethingFromB() {
  o := b.NewB()
  o.DoSomethingWithB()
}
