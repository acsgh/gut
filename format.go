package main

import (
	"github.com/albertoteloko/gutils/log"
	"os/exec"
)

func formatTask(dir string) {
	files, err := getFiles(dir, and(isGoFile, isGitHubPath))

	if err != nil {
		log.Error("Error during folder read: %s", err)
	}

	for _, f := range files {
		log.Debug("Formating file: %s", f)
		if err := formatFile(f); err != nil {
			log.Error("Error formating file %s, %s", f, err)
		} else {
			log.Debug("Formated file: %s", f)
		}
	}
}

func formatFile(file string) error {
	return exec.Command("gofmt", "-w", file).Run()
}
