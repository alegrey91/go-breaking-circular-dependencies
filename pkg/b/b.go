/* Previous b.go implementation, that return "import cycle not allowed" compilation error
package b

import (
    "fmt"
    "github.com/alegrey91/go-breaking-circular-dependencies/pkg/a"
)

type B struct {}

func NewB() *B {
    b := B{}
    return &b
}

func (b *B) DoSomethingWithB() {
    fmt.Println(b)
}

func InvokeSomethingFromA() {
    o := a.NewA()
    o.DoSomethingFromA()
}
*/
package b

import (
    "fmt"
)

type B struct {}

// This interface needs to be created in order to 
// allow pkg a to run its own method
type ADoSomethingIntf interface {
  DoSomethingWithA()
}

func NewB() *B {
    b := B{}
    return &b
}

func (b *B) DoSomethingWithB() {
  fmt.Println("This is b")
}

// This method needs to be modified (passing the interface as argument)
func InvokeSomethingFromA(o ADoSomethingIntf) {
  o.DoSomethingWithA()
}
