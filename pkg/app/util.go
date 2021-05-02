package app

import (
	strip "github.com/grokify/html-strip-tags-go"
	"regexp"
)

func CleanHTMLTags(content string) (imageURLs []string, result string) {
	var matches [][]string
	imageRex := regexp.MustCompile("<img.*?src=\"(.*?)\"(.*?)alt=\"(.*?)\"> ")
	matches = imageRex.FindAllStringSubmatch(content, -1)
	cleanedContent := imageRex.ReplaceAllString(content, "")

	for _, val := range matches {
		imageURLs = append(imageURLs, val[1])
	}

	result = strip.StripTags(cleanedContent)
	return
}
