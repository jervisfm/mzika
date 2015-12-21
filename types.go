package mzika

// Author: Jervis Muindi
// Date: December 2015

 
// NOTE: Names of the struct fields *MUST* start with an uppercase Letter so that
// they are considered "exported" and will be used when de-serializing the JSON
// messages
type VideoVersion struct {
	Version int
	SourceType int
	Id string
	// Don't index the data string because it exceeds 500 chars which is appengine
	// data store limit for indexed strings.
	// See: https://cloud.google.com/appengine/docs/go/datastore/reference
	Data string `datastore:",noindex"`
}

type Artist struct {
	ArtistId string
	ArtistName string
	ImageUrl string
}
type Video struct {
	// Video ID
	Isrc string
	Title string
	ImageUrl string
	DeepLinkUrl string
	IsPremiere bool
	Duration float64
	IsExplicit bool
	CopyrightLine string
	MainArtists []Artist
	VideoVersions []VideoVersion
}

type VideoJSON struct {
	CountryCode string
	LanguageCode string
	Video Video
}

type Rendition struct {
	XMLName xml.Name `xml:"rendition"`
	Name string `xml:"name,attr"`
	Url string `xml:"url,attr"`
	TotalBitrate int `xml:"totalBitrate,attr"`
	VideoBitrate int `xml:"videoBitrate,attr"`
	AudioBitrate int `xml:"audioBitrate,attr"`
	VideoCodec string `xml:"videoCodec,attr"`
	AudioCode string `xml:"audioCodec,attr"`
}

type Renditions struct {
	XMLName xml.Name `xml:"renditions"`
	Rendition []Rendition `xml:"rendition"`
}


// JSON Listing for top videos and searched videos
type VideoJSONListing struct {
	Success bool
    Result []struct {
      Isrc string
	  Title string
	  Image_url string
	}
}
