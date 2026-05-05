/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	RemoteURL     string `json:"remote_url"`
	DotfilesDir   string `json:"dotfiles_dir"`
	IncludeHidden bool   `json:"include_hidden"`
}

func defaultConfig(home string) Config {
	return Config{
		RemoteURL:   "",
		DotfilesDir: filepath.Join(home, ".dotfiles"),
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new dots configuration",
	Long:  "",
	Run:   setupConfigFile,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func setupConfigFile(cmd *cobra.Command, args []string) {
	// Check if the configuration file exists
	cfgPath := viper.ConfigFileUsed()
	if cfgPath != "" {
		// TODO: Check if the config file is valid (all necessary fields are set)
		fmt.Printf("You are all set, your config file is at %s%s%s\n", colorYellow, cfgPath, colorReset)
		fmt.Printf("To edit your config file, run %sdots config%s\n", colorBlue, colorReset)
		return
	}

	fmt.Printf("No config file found, creating at %s$HOME/.config/dots.json%s\n\n", colorYellow, colorReset)

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("%sOops! Error getting home directory: %v%s\n", colorRed, err, colorReset)
		return
	}
	cfgPath = filepath.Join(home, ".config", "dots.json")

	// Initialize new default configuration
	cfg := defaultConfig(home)
	reader := bufio.NewReader(os.Stdin)

	// Read remote URL from user input
	var remoteURL string
	fmt.Printf("Enter your remote dotfiles git URL: %s", colorGreen)

	remoteURL, err = reader.ReadString('\n')
	remoteURL = strings.TrimSpace(remoteURL)
	fmt.Printf("%s\n", colorReset)
	if err != nil {
		fmt.Printf("%sOops! Error reading remote URL: %v%s\n", colorRed, err, colorReset)
		return
	}

	// Parse remote URL and set config values
	parsedURL, err := url.Parse(remoteURL)
	if err != nil {
		fmt.Printf("%sOops! Error parsing remote URL: %v%s\n", colorRed, err, colorReset)
		return
	}

	if !strings.HasPrefix(parsedURL.String(), "https://") {
		parsedURL.Scheme = "https"
	}

	if !strings.HasSuffix(parsedURL.String(), ".git") {
		parsedURL.Path += ".git"
	}

	resp, err := http.Get(parsedURL.String())

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("%sOops! Error checking remote URL: %v%s\n", colorRed, err, colorReset)
		resp.Body.Close()
		return
	}

	defer resp.Body.Close()

	cfg.RemoteURL = parsedURL.String()

	// Prompt for the dotfiles directory
	var dotfilesDirInput string
	fmt.Printf("Enter the local dotfiles directory (default: %s): %s", cfg.DotfilesDir, colorGreen)

	dotfilesDirInput, err = reader.ReadString('\n')
	fmt.Printf("%s\n", colorReset)
	if err != nil {
		fmt.Printf("%sOops! Error reading dotfiles directory: %v%s\n", colorRed, err, colorReset)
		return
	}
	dotfilesDirInput = strings.TrimSpace(dotfilesDirInput)
	if dotfilesDirInput != "" {
		cfg.DotfilesDir = dotfilesDirInput
	}

	jsonCfg, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Printf("%sOops! Error marshaling config: %v%s\n", colorRed, err, colorReset)
		return
	}

	err = os.WriteFile(cfgPath, jsonCfg, 0644)
	if err != nil {
		fmt.Printf("%sOops! Error writing config: %v%s\n", colorRed, err, colorReset)
		return
	}

	fmt.Printf("All set! Run %sdots sync%s to sync your dotfiles.\n\n", colorBlue, colorReset)
}
