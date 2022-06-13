package test_utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var contentBodyRegexp = regexp.MustCompile("(?s)^---.+?---\n(.+)$")
var postFilePathAttributesRegexp = regexp.MustCompile("/(\\d{4})/([-a-z0-9]+)\\.md$")
var imageFilePathRegexp = regexp.MustCompile("!\\[[^]]*\\]\\(([^)]+)\\)")

// "github.com/gernest/front" で解析できる予定だったが、不具合で複数の "---" を含むときに正常に解析できなかった。
// Ref) https://github.com/gernest/front/pull/3
func ParseContentBody(content string) (contentBody string, err error) {
	matches := contentBodyRegexp.FindStringSubmatch(string(content))
	if len(matches) != 2 {
		return "", errors.New("The front matter and body cannot be separated.")
	}
	return matches[1], nil
}

func ParseImageFilePaths(contentBody string) []string {
	submatches := imageFilePathRegexp.FindAllStringSubmatch(contentBody, -1)
	paths := []string{}
	for _, submatch := range submatches {
		pathOrUrl := submatch[1]
		if !strings.HasPrefix(pathOrUrl, "http") {
			paths = append(paths, pathOrUrl)
		}
	}
	return paths
}

// 画像パスの "/{year}/{slug}/" と Markdown ファイルの "/{year}/{slug}.md" が一致するかを検証する。
func ValidateImageFilePath(imageFilePath string, postFilePath string) (bool, error) {
	postAttrsMatches := postFilePathAttributesRegexp.FindStringSubmatch(postFilePath)
	if len(postAttrsMatches) != 3 {
		return false, errors.New("This post file is in the wrong location.")
	}
	if strings.HasPrefix(imageFilePath, "http") {
		return true, nil
	}
	year := postAttrsMatches[1]
	urlSlug := postAttrsMatches[2]
	validImageFilePathPrefix := fmt.Sprintf("/external/posts/%s/%s/", year, urlSlug)
	return strings.HasPrefix(imageFilePath, validImageFilePathPrefix), nil
}
