package cmd

import "fmt"

const (
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorRed    = "\033[31m"
	colorReset  = "\033[0m"
)

const (
	fontBold  = "\033[1m"
	fontReset = "\033[0m"
)

func printError(msg string, err error) {
	fmt.Printf("%sOops! %s: %v%s\n", colorRed, msg, err, colorReset)
}
