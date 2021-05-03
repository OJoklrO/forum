package app

import (
	strip "github.com/grokify/html-strip-tags-go"
	"regexp"
	"time"
)

func CleanHTMLTags(content string) (imageURLs []string, result string) {
	var matches [][]string
	//imageRex := regexp.MustCompile("<img.*?src=\"(.*?)\"(.*?)alt=\"(.*?)\"> ")
	imageRex := regexp.MustCompile("<img.*?src=\"(.*?)\"(.*?)>")
	matches = imageRex.FindAllStringSubmatch(content, -1)
	cleanedContent := imageRex.ReplaceAllString(content, "")

	for _, val := range matches {
		imageURLs = append(imageURLs, val[1])
	}

	result = strip.StripTags(cleanedContent)
	return
}

func TimeFormat(unixTime int64) (timeStr string) {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	thisYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	commentTime := time.Unix(unixTime, 0)
	if today.Before(commentTime) {
		timeStr = commentTime.Format("15:04")
	} else if thisYear.Before(commentTime) {
		timeStr = commentTime.Format("01-02")
	} else {
		timeStr = commentTime.Format("2006-01-02")
	}
	return
}
