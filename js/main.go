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

// We wrap the exported calls to JS in go-routine so that 
// we avoid blocking. Output is returned via callback.

func GetVideoUrl(vid string, callback* js.Object) {
	go func() { 
		url, err := mzika.GetVideoUrl(vid)
		if callback != nil {
			callback.Invoke(url, err)
		}
	}()
}

func GetVideoFromId(vid string, callback* js.Object) {
	go func() {
		video_struct, err := mzika.GetVideoFromId(vid)
		if callback != nil {
			callback.Invoke(video_struct, err)
		}
	}()
}

func DecodeVideoJSON(input string, callback* js.Object) {
	go func() {
		video_struct, err := mzika.DecodeVideoJSON(input)
		if callback != nil {
			callback.Invoke(video_struct, err)
		}
	}()
}

func GetVideoRedirectUrl(input mzika.VideoJSON, callback* js.Object) {
	go func() {
		url, err := mzika.GetVideoRedirectUrl(input)
		if callback != nil {
			callback.Invoke(url, err)
		}
	}()
}

func main() {
	// Specify the method that we want to make available to Javascript
	// Our convention: all the method re-defined here are async by default (i.e. use goroutines). Methods under mzika.* can be blocking.
	js.Global.Set("mzika", map[string]interface{} {
		"decodeVideoJson": DecodeVideoJSON,
		"decodeVideoJsonSync": mzika.DecodeVideoJSON,

		"getVideoUrl" : GetVideoUrl,
		"getVideoFromId" : GetVideoFromId,

		// Added a sync version since no blocking i/o should be involved.
		"getVideoRedirectUrl" : GetVideoRedirectUrl,
		"getVideoRedirectUrlSync" : mzika.GetVideoRedirectUrl,
		
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
