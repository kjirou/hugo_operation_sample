package posts_test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gernest/front"

	"hugo_operation_sample/scripts/test_utils"
)

var fileNameRegexp = regexp.MustCompile("^[-a-z0-9]{1,64}\\.md$")
var dateRegexp = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$")

func validateFileName(filePath string) bool {
	_, fileName := filepath.Split(filePath)
	return fileNameRegexp.MatchString(fileName)
}

func validateDate(date string) bool {
	// time.Parse による検証のみだと "1:00:00" がエラーにならないため、
	//   正規表現による検証も維持している。
	if !dateRegexp.MatchString(date) {
		return false
	}
	if _, err := time.Parse("2006-01-02 15:04:05", date); err != nil {
		return false
	}
	return true
}

func TestPosts(t *testing.T) {
	testFileDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	projectRoot := filepath.Join(testFileDir, "..")

	type testCase struct {
		name     string
		filePath string
	}
	var testCases []testCase
	err = filepath.Walk(filepath.Join(projectRoot, "content/posts"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() == false && filepath.Ext(path) == ".md" {
			relPath, err := filepath.Rel(projectRoot, path)
			if err != nil {
				return err
			}
			testCases = append(testCases, testCase{
				name:     fmt.Sprintf("Post in %s", relPath),
				filePath: path,
			})
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !validateFileName(testCase.filePath) {
				t.Fatal("This is an invalid file name.")
			}

			content, err := os.ReadFile(testCase.filePath)
			if err != nil {
				t.Fatal(err)
			}

			frontMatter := front.NewMatter()
			frontMatter.Handle("---", front.YAMLHandler)
			fields, _, err := frontMatter.Parse(strings.NewReader(string(content)))
			if err != nil {
				t.Fatal(err)
			}

			if _, ok := fields["title"]; !ok {
				t.Fatal("There is no \"title\" field in the front matter.")
			}
			if date, ok := fields["date"]; !ok {
				t.Fatal("There is no \"date\" field in the front matter.")
			} else if dateAsString, ok := date.(string); !ok {
				t.Fatal("\"date\" is not a string.")
			} else if !validateDate(dateAsString) {
				t.Fatal("\"date\" format is invalid.")
			}
			if authors, ok := fields["authors"]; !ok {
				t.Fatal("There is no \"authors\" field in the front matter.")
			} else if authorsAsSlice, ok := authors.([]interface{}); !ok {
				t.Fatal("\"authors\" is not an array.")
			} else if len(authorsAsSlice) != 1 {
				t.Fatal("Set only one to \"authors\" field.")
			} else if _, ok := authorsAsSlice[0].(string); !ok {
				t.Fatal("The first of \"authors\" is not a string.")
			}

			contentBody, err := test_utils.ParseContentBody(string(content))
			if err != nil {
				t.Fatal(err)
			}

			for _, imageFilePath := range test_utils.ParseImageFilePaths(contentBody) {
				if isValid, err := test_utils.ValidateImageFilePath(imageFilePath, testCase.filePath); err != nil {
					t.Fatal(err)
				} else if !isValid {
					t.Fatalf("The location of the \"%s\" is incorrect.", imageFilePath)
				}
				realImageFilePath := filepath.Join(projectRoot, "static", imageFilePath)
				if fileInfo, err := os.Stat(realImageFilePath); os.IsNotExist(err) || fileInfo.IsDir() {
					t.Fatalf("\"%s\" does not exist for \"%s\".", realImageFilePath, imageFilePath)
				}
			}
		})
	}
}
