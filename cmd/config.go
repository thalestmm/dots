/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update the dots configuration",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfgPath := viper.ConfigFileUsed()

		if cfgPath == "" {
			initCmd.Run(cmd, args)
		}

		nvimCmd := exec.Command("nvim", cfgPath)
		nvimCmd.Stdin = os.Stdin
		nvimCmd.Stdout = os.Stdout
		nvimCmd.Stderr = os.Stderr

		if err := nvimCmd.Run(); err != nil {
			fmt.Printf("%sOops! An error occurred while opening the config file: %v%s\n", colorRed, err, colorReset)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
