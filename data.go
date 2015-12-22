package mzika

// Author: Jervis Muindi
// Date: December 2015

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Loads json listing of the top Music Videos into go structs |output|
func loadTopVideoJSONListing() (output VideoJSONListing, err error) {
	topVideosUrl := "https://api.vevo.com/mobile/v1/video/list.json?order=mostviewedthisweek&max=200"
	jsonContent, err := loadURL(topVideosUrl)
	if err != nil {
		err = fmt.Errorf("%v\n: Failed to fetch topvideoURL JSON", err)
	}
	return parseTopVideoJSONListing(jsonContent)
}

func loadSearchedVideoJSONListing(searchString string) (output VideoJSONListing, err error) {
	searchStringUrlEncoded := url.QueryEscape(searchString)
	searchUrlTemplate := "http://api.vevo.com/mobile/v1/search/videos.json?q=%s&max=30"
	url := fmt.Sprintf(searchUrlTemplate, searchStringUrlEncoded)
	jsonContent, _ := loadURL(url)
	if err != nil {
		err = fmt.Errorf("%v\n: Failed to fetch Searched JSON Url", err)
	}
	return parseTopVideoJSONListing(jsonContent)
}

// Takes given |vid| video identifier string and retrieves the JSON metadata associated with
// the specific video into |json|.
func loadVideoJSON(vid string) (json string, err error) {
	videoJsonURLTemplate := "http://videoplayer.vevo.com/VideoService/AuthenticateVideo?isrc=%s"
	url := fmt.Sprintf(videoJsonURLTemplate, vid)
	return loadURL(url)
}

// Loads given |url| string and returns a |response| with the output. |url| should contain
// an appropriate protocol (e.g "http://www.msn.com")
func loadURL(url string) (response string, err error) {
	println("Loading url: ", url)
	resp, err := http.Get(url)
	if err != nil {
		println("url fail")
		return "", fmt.Errorf("Failed to fetch URL: %s. Got Response : %s", url, resp)
	}
	// Ensure that we close the reading handle upon function exit.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to Read All Data from URL: %s. Got Response Code: %s.", url, resp.Status)
	}
	println(string(body))
	return string(body), err
}

func parseTopVideoJSONListing(input string) (output VideoJSONListing, err error) {
	var m VideoJSONListing
	err = json.Unmarshal([]byte(input), &m)
	if err != nil {
		return m, fmt.Errorf("Failed to Decode Json: %s. \n==Error:'%s'", input, err)
	}
	return m, err
}

// Attempts to decode given |input| JSON string into a VideoJSON go struct.
func DecodeVideoJSON(input string) (output VideoJSON, err error) {
	var m VideoJSON
	err = json.Unmarshal([]byte(input), &m)
	if err != nil {
		return m, fmt.Errorf("Failed to Decode Json: %s. \n==Error:'%s'", input, err)
	}
	return m, err
}

// Attempts to decode given |input| XML string into a Renditions go struct type.
func DecodeVideoRendition(input string) (output Renditions, err error) {
	err = xml.Unmarshal([]byte(input), &output)
	if err != nil {
		return output, fmt.Errorf("Failed to Decode XML: %s. \n==Error:'%s'", input, err)
	}
	return output, err
}
