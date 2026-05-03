// Copyright 2026 @thalestmm. All rights reserved.
//
// MIT license
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
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
	gitRemoteURLInput := flag.String("git", "", "URL for the git remote repository")
	targetDir := flag.String("dir", ".", "relative path to target directory")
	includeHidden := flag.Bool("hidden", false, "include hidden directories in target directory (nested hidden directories are always included)")
	isDryRun := flag.Bool("dry-run", false, "don't apply symlinks")
	flag.Parse()

	var gitRemoteURL *url.URL
	if *gitRemoteURLInput != "" {
		fmt.Println("Parsing git URL...")
		gitRemoteURL, err := url.Parse(*gitRemoteURLInput)
		if err != nil {
			fmt.Printf("%sOops! Failed to parse the git URL: %v%s\n", colorRed, err, colorReset)
			os.Exit(1)
		}
		fmt.Printf("\nChecked! %s%s%s\n", colorYellow, gitRemoteURL, colorReset)

		// Check if git is installed
		if err := exec.Command("git", "-v").Start(); err != nil {
			fmt.Printf("%sOops! git is not installed.%s\n", colorRed, colorReset)
			os.Exit(1)
		}
	}

	fmt.Println(gitRemoteURL)

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

	fmt.Println()

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

	fmt.Printf("Upserted .dotfiles directory: %s%s%s\n", colorYellow, dotfilesDir, colorReset)

	// Copy contents of target directory to .dotfiles dir
	for _, entry := range dirs {
		if err := copyDir(filepath.Join(wd, entry.Name()), dotfilesDir); err != nil {
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

	var dotfileDirs []os.DirEntry
	for _, entry := range dirs {
		if entry.IsDir() {
			fmt.Printf("> %s%s%s\n", colorYellow, entry.Name(), colorReset)
			dotfileDirs = append(dotfileDirs, entry)
		}
	}

	fmt.Println()

	// Traverse each .dotfiles directory and symlink to the desired path

	fmt.Printf("Starting synchronization\n\n")
	for _, dir := range dotfileDirs {
		// Skip git dir
		if dir.Name() != ".git" {
			fmt.Printf("Working on %s%s%s... ", colorGreen, dir.Name(), colorReset)
			if !*isDryRun {
				if err := traverse(filepath.Join(dotfilesDir, dir.Name()), "", homeDir); err != nil {
					fmt.Printf("%sOops! Failed to traverse and symlink: %v%s\n", colorRed, err, colorReset)
					os.Exit(1)
				}
			}
			fmt.Printf("Done!\n\n")
		}
	}

	fmt.Println()

	os.Exit(0)
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
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	dirName := filepath.Base(src)

	if err := os.MkdirAll(filepath.Join(dst, dirName), 0755); err != nil {
		return err
	}

	for _, entry := range entries {
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

	return nil
}

// traverse walks recursively trough a dotfiles directory and symlink
// the children files to each mapped path.
func traverse(path, dirName, homeDir string) error {
	if err := os.MkdirAll(filepath.Join(homeDir, dirName), 0755); err != nil {
		return err
	}
	children, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, child := range children {
		srcPath := filepath.Join(path, child.Name())
		dstPath := filepath.Join(homeDir, dirName, child.Name())

		if child.IsDir() {
			if err := traverse(filepath.Join(path, child.Name()), filepath.Join(dirName, child.Name()), homeDir); err != nil {
				return err
			}
		} else {
			if err := exec.Command("ln", "-s", srcPath, dstPath).Start(); err != nil {
				return err
			}
		}
	}
	return nil
}
