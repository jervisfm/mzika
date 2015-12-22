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
