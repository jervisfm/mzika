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
	// Specify the method that we want to make available to Javascript
	js.Global.Set("mzika", map[string]interface{} {
		"decodeVideoJson": mzika.DecodeVideoJSON,
		"getVideoUrl" : mzika.GetVideoUrl,
		"getVideoFromId" : mzika.GetVideoFromId,
	})
	fmt.Println("Hello, playground")
	//println("Hello, JS console")
}
