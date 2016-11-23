package main

import (
	"github.com/albertoteloko/gutils/log"
	"strings"
	"os/exec"
)

func benchmarkTask(dir string) {
	files, err := getFolders(dir, and(isGitHubPath, isTestFolder))

	if err != nil {
		log.Error("Error during folder read: %s", err)
	}

	for _, f := range files {
		log.Info("Benchmarking golder: %s", f)

		results, err := benchmarkFolder(f)

		if err != nil {
			log.Error("Error during folder benchmarking: %s", err)
		} else {
			namesLength := 0;
			for _, result := range results {
				namesLength = max(namesLength, len(result.name))
			}
			for _, result := range results {
				log.Info("%s %s [%s]", fixWidth(result.name, namesLength + 3), result.rate, result.times)
			}
		}

		log.Debug("Benchmarked folder: %s", f)
	}
}



type benchmarkResult struct {
	name  string
	times string
	rate  string
}

func benchmarkFolder(folder string) ([]benchmarkResult, error) {
	command := exec.Command("go", "test", folder, "-bench", ".", "-run", "NoOneWillFitThisDescription")
	bytes, err := command.Output()

	output := string(bytes)

	if !strings.Contains(output, "FAIL") && err != nil {
		return []benchmarkResult{}, err
	} else {
		testOutputs := []benchmarkResult{}

		for _, line := range strings.Split(output, "\n") {
			if strings.HasPrefix(line, "Benchmark") {
				parts := strings.Split(replaceChars(line, "", " ", "---", ")"), "\t")
				testOutputs = append(testOutputs, benchmarkResult{parts[0], parts[1], parts[2]})
			}
		}

		return testOutputs, nil
	}
}
