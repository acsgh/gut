package main

import (
	"flag"
	"github.com/albertoteloko/gutils/log"
	"os"
	"time"
)

const version = "1.0"

var directory string

var debug bool
var recursive bool

var (
	GO_PATH = os.Getenv("GOPATH")
	BIN_FOLDER = GO_PATH + "/bin"
)

func main() {
	var clean bool
	var build bool
	var test bool
	var bench bool
	var format bool
	var all bool

	flag.BoolVar(&debug, "v", false, "Vervose Output")
	flag.BoolVar(&recursive, "r", false, "Recursive")

	flag.BoolVar(&clean, "c", false, "Execute Clean Task")
	flag.BoolVar(&build, "b", false, "Execute Build Task")
	flag.BoolVar(&test, "t", false, "Execute Test Task")
	flag.BoolVar(&bench, "bench", false, "Execute Benchmark Task")
	flag.BoolVar(&format, "f", false, "Execute Format Task")

	flag.BoolVar(&all, "a", false, "Execute All Task")

	flag.StringVar(&directory, "dir", "", "Base Directory")
	flag.Parse()

	if debug {
		log.Level(log.DEBUG)
	} else {
		log.Level(log.INFO)
	}

	log.Info("GUT %s: Start", version)
	log.Debug("GOPATH: %s", GO_PATH)

	//tasks := flag.Args()
	startTime := time.Now().UnixNano()

	dir := getBaseDir()
	log.Debug("Base Dir: %s", dir)

	if clean || all {
		logTime("clean", func() {
			cleanTask()
		})
	}

	if format || all {
		logTime("format", func() {
			formatTask(dir)
		})
	}

	if build || all {
		logTime("build", func() {
			buildTask(dir)
		})
	}

	allTestPass := true
	if test || all {
		logTime("test", func() {
			allTestPass = testTask(dir)
		})
	}

	if (!allTestPass) {
		log.Fatal("Some test fails")
	} else {
		if bench || all {
			logTime("benchmark", func() {
				benchmarkTask(dir)
			})
		}
	}
	log.Info("All Tasks: Done in %d ms", (time.Now().UnixNano() - startTime) / 1000000)

	//for i := 0; i < len(tasks); i++ {
	//	task := tasks[i]
	//	switch task {
	//	case "clean":
	//		logTime("clean", clean)
	//	case "build":
	//		logTime("build", build)
	//	case "format":
	//		logTime("format", format)
	//	case "test":
	//		logTime("test", test)
	//	case "get":
	//		logTime("get", get)
	//	default:
	//		log.Error("Task not found: %s", task)
	//	}
	//}
}

func logTime(taskName string, task Task) {
	startTime := time.Now().UnixNano()
	log.Info("Task %s: Start", taskName)
	task()
	log.Info("Task %s: Done in %d ms", taskName, (time.Now().UnixNano() - startTime) / 1000000)
}

func build() {
}
func test() {
}
func get() {
}

type Task func()
