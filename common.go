package main

import (
	"io/ioutil"
	"strings"
)

func replaceChars(input string, replaceBy string, values ...string) string {
	result := input

	for _, value := range values {
		result = strings.Replace(result, value, replaceBy, -1)
	}

	return result
}

func max(v1, v2 int) int {
	if (v1 >= v2) {
		return v1
	} else {
		return v2
	}
}

func fixWidth(value string, width int) string {
	result := value

	for len(result) < width {
		result += " "
	}

	return result
}

func and(f1, f2 func(string) bool) func(string) bool {
	return func(file string) bool {
		return f1(file) && f2(file)
	}
}

func isGoFile(file string) bool {
	return strings.HasSuffix(file, ".go")
}

func isGoTestFile(file string) bool {
	return strings.HasSuffix(file, "_test.go")
}

func isGitHubPath(file string) bool {
	return strings.Contains(file, "github.com/albertoteloko")
}

func getFiles(name string, filter func(string) bool) ([]string, error) {
	result := []string{}

	files, err := ioutil.ReadDir(name)

	if err != nil {
		return result, err
	}

	for _, f := range files {
		fileName := name + "/" + f.Name()
		if recursive && f.IsDir() {
			childFiles, err := getFiles(fileName, filter)

			if err != nil {
				return result, err
			}
			result = append(result, childFiles...)

		} else if filter(fileName) {
			result = append(result, fileName)
		}
	}

	return result, nil
}

func getFolders(name string, filter func(string) bool) ([]string, error) {
	result := []string{}

	files, err := ioutil.ReadDir(name)

	if err != nil {
		return result, err
	}

	for _, f := range files {
		fileName := name + "/" + f.Name()
		if f.IsDir() {
			if recursive {
				childFolders, err := getFolders(fileName, filter)

				if err != nil {
					return result, err
				}
				result = append(result, childFolders...)
			}
			if filter(fileName) {
				result = append(result, fileName)
			}
		}
	}

	return result, nil
}

func getBaseDir() string {
	if directory != "" {
		return directory
	} else {
		return "."
	}
}
