package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	targetDir := flag.String("dir", ".", "target directory")
	includeHidden := flag.Bool("hidden", false, "include hidden directories")
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

	fmt.Print("Directories found:\n\n")

	for _, entry := range contents {
		if entry.IsDir() {
			if !*includeHidden && entry.Name()[0] == '.' {
				continue
			}
			fmt.Println(entry.Name())
		}
	}
}
