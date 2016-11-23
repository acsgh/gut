package main

import (
	"os"
)

func cleanTask() {
	os.RemoveAll(BIN_FOLDER)
	os.MkdirAll(BIN_FOLDER, 0777)
}
