/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

const minPerfectMatchLength = 20
const minPossibleMatchLength = 10

type MatchDetails struct {
	MatchType   string
	MatchFile   string
	MatchedLine string
	MatchedWord string
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "The command will recursive scan directories for possible bitcoin addresses",
	Long:  `The command will recursive scan directories for possible bitcoin addresses.`,
	Run: func(cmd *cobra.Command, args []string) {

		scanDirectoryPath := ""
		if len(args) >= 1 && args[0] != "" {
			scanDirectoryPath = args[0]
		} else {
			fmt.Println("A valid directory path is required")
			return
		}

		if matchedFileList, err := ScanDirectory(scanDirectoryPath); err != nil {
			fmt.Printf("Scanning Failed for directory: %s err: %s\n", scanDirectoryPath, err.Error())
		} else {
			for _, matchedDetails := range matchedFileList {
				if prettyJson, err := json.MarshalIndent(*matchedDetails, "", "    "); err != nil {
					fmt.Printf("Failed to jsonify object: %+v err: %s\n", *matchedDetails, err.Error())
				} else {
					fmt.Printf("%s\n", string(prettyJson))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func ScanDirectory(scanDirectoryPath string) ([]*MatchDetails, error) {

	if _, err := os.Stat(scanDirectoryPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("No directory found at " + scanDirectoryPath)
	}

	matchDetailsList := []*MatchDetails{}

	err := filepath.Walk(scanDirectoryPath, func(filePath string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			if fileMatchList, err := ScanFile(filePath); err != nil {
				return err
			} else {
				matchDetailsList = append(matchDetailsList, fileMatchList...)
			}
		}
		return nil
	})

	return matchDetailsList, err
}

func ScanFile(scanFilePath string) ([]*MatchDetails, error) {
	if scanFileInfo, err := os.Stat(scanFilePath); os.IsNotExist(err) || scanFileInfo.IsDir() {
		return nil, fmt.Errorf("No file found at " + scanFilePath)
	}

	scanFile, err := os.Open(scanFilePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: " + scanFilePath + " err: " + err.Error())
	}
	defer scanFile.Close()

	matchDetailsList := []*MatchDetails{}

	scanner := bufio.NewScanner(scanFile)
	base58Re := regexp.MustCompile(`[a-km-zA-HJ-NP-Z1-9]+`)
	base58Re.Longest()
	for scanner.Scan() {
		fileLine := scanner.Text()
		if len(fileLine) > 0 {
			for _, word := range strings.Split(fileLine, " ") {
				matchedString := findLongestMatch(base58Re, word)
				matchedStringLength := len(matchedString)
				if len(word) > minPerfectMatchLength && matchedStringLength > minPossibleMatchLength {
					if matchedStringLength == len(word) {
						matchDetailsList = append(matchDetailsList, &MatchDetails{MatchType: "perfect", MatchFile: scanFilePath, MatchedLine: fileLine, MatchedWord: word})
					} else {
						matchDetailsList = append(matchDetailsList, &MatchDetails{MatchType: "possible", MatchFile: scanFilePath, MatchedLine: fileLine, MatchedWord: word})
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Occurred while scanning file: " + scanFilePath + "err: " + err.Error())
		return nil, err
	}

	return matchDetailsList, nil
}

func findLongestMatch(regex *regexp.Regexp, matchStr string) string {
	matchStrings := regex.FindAllString(matchStr, -1)
	longestMatch := ""
	for _, matchString := range matchStrings {
		if len(matchString) > len(longestMatch) {
			longestMatch = matchString
		}
	}

	return longestMatch
}
