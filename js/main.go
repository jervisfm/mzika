// This file is meant to expose useful functions in Mzika to be available
// in Javascript.
// We're putting the code in the main package because otherwise
// gopherjs refuses to generate the (transpiled) Javascript code.
package main
import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/jervisfm/mzika"
)

func main() {
	fmt.Println("Hello, playground")
	_,_ = mzika.DecodeVideoJSON("")
	js.Global.Call("alert", "Hello, JavaScript")
	println("Hello, JS console")
}
