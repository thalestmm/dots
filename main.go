package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorRed    = "\033[31m"
	colorReset  = "\033[0m"
)

func main() {
	targetDir := flag.String("dir", ".", "target directory")
	includeHidden := flag.Bool("hidden", false, "include hidden directories")
	flag.Parse()

	err := os.Chdir(*targetDir)
	if err != nil {
		fmt.Printf("%sOops! Failed to change directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}

	wd, err := os.Getwd()

	fmt.Printf("\nProcessing directory: %s%s%s\n\n", colorGreen, wd, colorReset)

	if err != nil {
		fmt.Printf("%sOops! Failed to get current directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}

	contents, err := os.ReadDir(wd)
	if err != nil {
		fmt.Printf("%sOops! Failed to read directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}

	fmt.Print("Directories found:\n\n")

	for _, entry := range contents {
		if entry.IsDir() {
			if !*includeHidden && entry.Name()[0] == '.' {
				continue
			}
			fmt.Printf("> %s%s%s\n", colorBlue, entry.Name(), colorReset)
		}
	}
}
