package mzika

// Author: Jervis Muindi
// Date: December 2015

// Loads json listing of the top Music Videos into go structs |output|
func loadTopVideoJSONListing(r *http.Request) (output VideoJSONListing, err error) {
	topVideosUrl := "https://api.vevo.com/mobile/v1/video/list.json?order=mostviewedthisweek&max=200"
	jsonContent, err := loadURL(r, topVideosUrl)
	if err != nil {
		err = fmt.Errorf("%v\n: Failed to fetch topvideoURL JSON", err)
	}
	return parseTopVideoJSONListing(jsonContent)
}

func loadSearchedVideoJSONListing(r *http.Request, searchString string) (output VideoJSONListing, err error) {
	searchStringUrlEncoded := url.QueryEscape(searchString)
	searchUrlTemplate := "http://api.vevo.com/mobile/v1/search/videos.json?q=%s&max=30"
	url := fmt.Sprintf(searchUrlTemplate, searchStringUrlEncoded)
	jsonContent, _ := loadURL(r, url)
	if err != nil {
		err = fmt.Errorf("%v\n: Failed to fetch Searched JSON Url", err)
	}
	return parseTopVideoJSONListing(jsonContent)
}

// Takes given |vid| video identifier string and retrieves the JSON metadata associated with
// the specific video into |json|.
func loadVideoJSON(r *http.Request, vid string) (json string, err error) {
	videoJsonURLTemplate := "http://videoplayer.vevo.com/VideoService/AuthenticateVideo?isrc=%s"
	url := fmt.Sprintf(videoJsonURLTemplate, vid)
	return loadURL(r, url)
}
