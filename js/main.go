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

func GetVideoUrl(vid string) {
	go func() { 
		url,err := mzika.GetVideoUrl(vid)
		println("Resolved Video URL:", url, err)
	}()
}

func main() {
	// Specify the method that we want to make available to Javascript
	js.Global.Set("mzika", map[string]interface{} {
		"decodeVideoJson": mzika.DecodeVideoJSON,
		//"getVideoUrl" : mzika.GetVideoUrl,
		"getVideoUrl" : GetVideoUrl,
		"getVideoFromId" : mzika.GetVideoFromId,
	})
	fmt.Println("Hello, playground")

	// Go test code snippet
	/* 
	vid := "uscmv1500002"
	url, err := mzika.GetVideoUrl(vid)
	if err != nil {
		println("Failure!", err)
		return
	}
	fmt.Println(url) */
}
