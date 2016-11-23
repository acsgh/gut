package main

import (
	"github.com/albertoteloko/gutils/log"
	"io/ioutil"
	"os/exec"
	"strings"
)

func testTask(dir string) bool {
	files, err := getFolders(dir, and(isGitHubPath, isTestFolder))

	if err != nil {
		log.Error("Error during folder read: %s", err)
		return false
	}

	result:= true
	for _, f := range files {
		log.Info("Testing golder: %s", f)

		results, coverage, err := testFolder(f)

		if err != nil {
			log.Error("Error during folder testing: %s", err)
		} else {
			namesLength := 0;
			for _, result := range results {
				namesLength = max(namesLength, len(result.name))
			}
			for _, testResult := range results {
				result = result && (testResult.result == "PASS")
				log.Info("%s %s [%s]", fixWidth(testResult.name, namesLength + 3), testResult.result, testResult.time)
			}
			log.Info("Coverage: %s", coverage)
		}

		log.Debug("Tested folder: %s", f)
	}
	return result
}

type testResult struct {
	name   string
	result string
	time   string
}

func testFolder(folder string) ([]testResult, string, error) {
	command := exec.Command("go", "test", folder, "-v", "-cover")
	bytes, err := command.Output()

	output := string(bytes)

	if !strings.Contains(output, "FAIL") && err != nil {
		return []testResult{}, "", err
	} else {
		testOutputs := []testResult{}
		coverage := "0.0%"
		for _, line := range strings.Split(output, "\n") {
			if strings.HasPrefix(line, "--- ") {
				parts := strings.Split(replaceChars(replaceChars(line, "", " ", "---", ")"), "\t", ":", "("), "\t")
				testOutputs = append(testOutputs, testResult{parts[1], parts[0], parts[2]})
			} else if strings.HasPrefix(line, "coverage: ") {
				coverage = replaceChars(line, "", "coverage: ", " of statements")
			}

		}

		return testOutputs, coverage, nil
	}
}

func isTestFolder(name string) bool {
	files, err := ioutil.ReadDir(name)

	if err != nil {
		return false
	}

	for _, f := range files {
		fileName := name + "/" + f.Name()
		if !f.IsDir() && isGoTestFile(fileName) {
			return true
		}
	}
	return false
}
