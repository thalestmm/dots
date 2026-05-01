package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	targetDir := flag.String("dir", ".", "target directory")
	flag.Parse()

	err := os.Chdir(*targetDir)
	if err != nil {
		fmt.Printf("Oops! Failed to change directory: %v\n", err)
		os.Exit(1)
	}

	wd, err := os.Getwd()

	if err != nil {
		fmt.Printf("Oops! Failed to get current directory: %v\n", err)
		os.Exit(1)
	}

	contents, err := os.ReadDir(wd)
	if err != nil {
		fmt.Printf("Oops! Failed to read directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(contents)
}
