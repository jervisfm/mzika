package mzika

// Author: Jervis Muindi
// Date: December 2015

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)


// Parses given width/height specification along with VideoJson object |videoJson|
// and returns a URL to an image thumbnail of the video. width/height can be
// both be set to 0 to recover original full-resolution image.
func GetImageUrl(width, height int, videoJson VideoJSON) string {
	url := videoJson.Video.ImageUrl
	// Add Height/Width params iff they are non-zero.
	url += "?"
	addAmpersand := false
	if width > 0 {
		url += fmt.Sprintf("width=%s", width)
		addAmpersand = true
	}
	if height > 0 {
		if addAmpersand {
			url += "&"
		}
		url += fmt.Sprintf("height=%s", height)
	}
	return url
}

// Returns a video url to watch a video with given |vid|. |vid| is a
// video Id that looks like "uscmv1500002".
func GetVideoUrl(vid string) (url string, err error) {
	v, err := GetVideoFromId(vid)
	if err != nil {
		fmt.Errorf("Failed to retrieve video %v. Error: %v", vid, err.Error())
		return
	}

	// Look for and Pick a viable URL to watch the video.
	url, err = GetVideoRedirectUrl(v)
	if err != nil {
		fmt.Errorf("Failed to find watchable video url for vid %v. Error: %v", vid, err.Error())
		return
	}
	return url,nil
}

// Takes given |vid| Video identifier and converts it into a VideoJSON struct containing
// appropriate metadata for the specific video.
func GetVideoFromId(vid string) (output VideoJSON, err error) {
	// Try to get data from data store cache
	vid = strings.ToLower(vid)
	cacheHit, cacheResponse, err := GetCachedVideoResponse(vid)
	if err == nil {
		if cacheHit {
			println("Gota cache hit, ", cacheResponse)
			output = *cacheResponse
			err = nil
			return
		}
	}

	// If we get here, it's for the cases:
	// 1) the response has not been cached yet, OR
	// 2) looking up the cached response failed.
	// In either case, we need to do the JSON fetch from the
	// interwebs and then cache this response.

	// Fetch Video JSON
	//
	// TODO(jervis):
	// Consider using HTML5 Persistent localstorage as a cache so that
	// we don't repeatedly look up the same URLs over and over across sessions.
	resp, err := loadVideoJSON(vid)
	if err != nil {
		return output, err
	}

	// Parse Into a VideoJSON struct
	output, err = DecodeVideoJSON(resp)
	// Cache the response so future looks are faster and avoid
	// network requests.
	cacheSaveError := CacheVideoJsonResponse(output, vid)
	// Only take note of cache save error but don't do anything else.
	if cacheSaveError != nil {
		fmt.Printf(">> Error:%v\nOops, failed to save cached response for video: '%s'", cacheSaveError, vid)
	}
	return output, err
}

// Attempts to parse given |input| of a VideoJSON struct
// and extract out a viable URL from which the video can be watched. By default, we try to
// pick the highest available MP4 stream from either AWS or LVL3 (Level3).
func GetVideoRedirectUrl(input VideoJSON) (output string, err error) {
	const kDefaultRedirect = "http://www.google.com"
	var awsVideoUrl *string = nil
	var level3VideoUrl *string = nil
	for _, video := range input.Video.VideoVersions {
		renditions, err := DecodeVideoRendition(video.Data)
		if err != nil {
			continue
		}
		for _, rendition := range renditions.Rendition {
			videoQuality := rendition.Name
			url := rendition.Url
			// Only interested in High Quality URLs of Mp4 videos
			if videoQuality == "High" && strings.Contains(url, ".mp4") {
				// And Further Limit to Amazon/Level3 Hosted URLs.
				const level3 = "lvl3"
				const amazon = "aws"
				if strings.Contains(url, level3) {
					level3VideoUrl = &url
				}
				if strings.Contains(url, amazon) {
					awsVideoUrl = &url
				}
			}
		}
	}
	// Prefer aws url as it's more stable.
	if awsVideoUrl != nil {
		return *awsVideoUrl, nil
	}
	if level3VideoUrl != nil {
		return *level3VideoUrl, nil
	}
	return kDefaultRedirect,
		fmt.Errorf("Failed to find a suitable HighQuality Amazon/Lvl3 based video URL for Video %s", input)
}

// Renders the Home page which lists the current most popular videos.
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	videoListing, err := LoadTopVideoJSONListing(DefaultOrder, FirstPage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Add templates/* to search path.
	var tpl = template.Must(template.ParseGlob("templates/*.html"))
	if err := tpl.ExecuteTemplate(w, "index.html", videoListing); err != nil {
		fmt.Errorf("%v", err)
	}
}

// Renders the Home page which lists the current most popular videos.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	videoListing, err := LoadSearchedVideoJSONListing(searchQuery, FirstPage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := struct {
		Listing     VideoJSONListing
		SearchQuery string
	}{videoListing, searchQuery}
	var tpl = template.Must(template.ParseGlob("templates/*.html"))
	if err := tpl.ExecuteTemplate(w, "search.html", data); err != nil {
		fmt.Errorf("%v", err)
	}
}
