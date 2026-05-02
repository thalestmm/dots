package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorRed    = "\033[31m"
	colorReset  = "\033[0m"
)

func main() {
	targetDir := flag.String("dir", ".", "relative path to target directory")
	includeHidden := flag.Bool("hidden", false, "include hidden directories in target directory (nested hidden directories are always included)")
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

	var dirs []os.DirEntry

	for _, entry := range contents {
		if entry.IsDir() {
			if !*includeHidden && entry.Name()[0] == '.' {
				continue
			}
			fmt.Printf("> %s%s%s\n", colorBlue, entry.Name(), colorReset)
			dirs = append(dirs, entry)
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("%sOops! Failed to get home directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}

	// See if the .dotfiles dir already exists
	dotfilesDir := homeDir + "/.dotfiles"

	if err := os.MkdirAll(dotfilesDir, 0755); err != nil {
		fmt.Printf("%sOops! Failed to create .dotfiles directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}

	fmt.Printf("\nUpserted .dotfiles directory: %s%s%s\n", colorYellow, dotfilesDir, colorReset)

	// Copy contents of target directory to .dotfiles dir
	for _, entry := range dirs {
		if err := copyDir(filepath.Join(wd, entry.Name()), dotfilesDir); err != nil {
			fmt.Printf("%s", filepath.Join(*targetDir, entry.Name()))
			fmt.Printf("%sOops! Failed to copy directory: %v%s\n", colorRed, err, colorReset)
			os.Exit(1)
		}
	}

	dirs, err = os.ReadDir(dotfilesDir)
	if err != nil {
		fmt.Printf("%sOops! Failed to read .dotfiles directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}

	fmt.Printf("\nContents of .dotfiles:\n\n")
	for _, entry := range dirs {
		if entry.IsDir() {
			fmt.Printf("> %s%s%s\n", colorYellow, entry.Name(), colorReset)
		}
	}

	// Traverse each .dotfiles directory and symlink to the desired path

	// TODO: Remove, debug only
	if err := exec.Command("open", dotfilesDir).Start(); err != nil {
		fmt.Printf("%sOops! Failed to open directory: %v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}
}

func copyFile(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, bytes, 0644); err != nil {
		return err
	}

	return nil
}

// copyDir copies the contents of a directory to another directory.
// It takes a full src path and the base dst path.
func copyDir(src, dst string) error {
	fmt.Printf("src: %s\n", src)
	fmt.Printf("dst: %s\n", dst)
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	dirName := filepath.Base(src)
	srcParentPath := filepath.Dir(src)
	fmt.Printf("parent: %s\n", srcParentPath)
	fmt.Printf("name: %s\n\n", dirName)

	if err := os.MkdirAll(filepath.Join(dst, dirName), 0755); err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Printf("entry: %s\n", entry.Name())
		if entry.IsDir() {
			if err := copyDir(filepath.Join(src, entry.Name()), filepath.Join(dst, dirName)); err != nil {
				return err
			}
		} else {
			if err := copyFile(filepath.Join(src, entry.Name()), filepath.Join(dst, dirName, entry.Name())); err != nil {
				return err
			}
		}
	}

	fmt.Println("---")
	return nil
}
