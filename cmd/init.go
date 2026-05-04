/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	RemoteURL   string `json:"remote_url"`
	DotfilesDir string `json:"dotfiles_dir"`
}

func defaultConfig() Config {
	return Config{
		RemoteURL:   "",
		DotfilesDir: "~/.dotfiles",
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new dots configuration",
	Long:  "",
	Run:   initializeConfiguration,
}

func initializeConfiguration(cmd *cobra.Command, args []string) {
	// Check if the configuration file exists
	cfgPath := viper.ConfigFileUsed()
	if cfgPath != "" {
		fmt.Println("You are all set, your config file is at", cfgPath)
		fmt.Println("To edit your config file, run \033[dots config`")
		return
	}

	fmt.Printf("No config file found, creating at %s$HOME/.config/dots.json%s\n\n", colorYellow, colorReset)

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	cfgPath = filepath.Join(home, ".config", "dots.json")

	var remoteURL string
	fmt.Printf("Enter your remote git URL: %s", colorGreen)
	_, err = fmt.Scanln(&remoteURL)
	fmt.Printf("%s\n", colorReset)
	if err != nil {
		fmt.Printf("%sOops! Error reading remote URL: %v%s\n", colorRed, err, colorReset)
		return
	}

	fmt.Println(remoteURL)

}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
