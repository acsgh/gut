package main

import (
	"os"
	"fmt"
	"flag"
)

func main() {
	args := os.Args[1:]
	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("tail:", flag.Args())

	fmt.Println(args)
}
