/*
Copyright © 2026 Thales Monteiro Meier @thalestmm

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/thalestmm/dots/cmd"
)

//go:embed app.json
var appInfoFile []byte

type AppInfo struct {
	Version string `json:"version"`
}

func main() {
	var appInfo AppInfo

	// Parse app.json
	if err := json.Unmarshal(appInfoFile, &appInfo); err != nil {
		fmt.Printf("Oops! Error parsing app.json: %v\n", err)
		os.Exit(1)
	}

	// Set version here to retrieve later
	viper.Set("version", appInfo.Version)

	cmd.Execute()
}
