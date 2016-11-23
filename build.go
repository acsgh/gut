package main

import (
	"github.com/albertoteloko/gutils/log"
	"io/ioutil"
	"os/exec"
	"strings"
)

func buildTask(dir string) bool {
	files, err := getFolders(dir, and(isGitHubPath, isGoFolder))

	//build -o "bin/$1" "github.com/albertoteloko/$1"

	if err != nil {
		log.Error("Error during folder read: %s", err)
		return false
	}

	result := true
	for _, f := range files {
		log.Info("Building golder: %s", f)

		_, err := buildFolder(f)

		if err != nil {
			log.Error("Error during folder building: %s", err)
		}

		log.Debug("Builded folder: %s", f)
	}
	return result
}

func getFolderName(folderPath string) string {
	if (strings.Contains(folderPath, "/")) {
		return folderPath[strings.LastIndex(folderPath, "/") + 1:]
	} else {
		return folderPath
	}
}

func buildFolder(folder string) (string, error) {
	command := exec.Command("go", "build", "-o", GO_PATH + "/bin/" + getFolderName(folder), folder)
	bytes, err := command.Output()
	return string(bytes), err
}

func isGoFolder(name string) bool {
	files, err := ioutil.ReadDir(name)

	if err != nil {
		return false
	}

	for _, f := range files {
		fileName := name + "/" + f.Name()
		if !f.IsDir() && isGoMainFile(fileName) {
			return true
		}
	}
	return false
}

func isGoMainFile(name string) bool {
	b, err := ioutil.ReadFile(name) // just pass the file name
	if err != nil {
		log.Error("Unable to read file %s: %s", name, err)
		return false
	}

	return strings.HasPrefix(strings.Trim(string(b), " \n\r"), "package main")
}
