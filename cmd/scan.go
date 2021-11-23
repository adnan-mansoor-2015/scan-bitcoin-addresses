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
	"time"

	"github.com/spf13/cobra"
)

const minPerfectMatchLength = 20
const minPossibleMatchLength = 10

type MatchDetails struct {
	MatchType     string
	MatchFile     string
	MatchedLine   string
	MatchedLineNo int
	MatchedWord   string
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "The command will recursive scan directories for possible bitcoin addresses",
	Long:  `The command will recursive scan directories for possible bitcoin addresses.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		directories := []string{}
		if len(args) >= 1 {
			for _, arg := range args {
				if len(arg) > 0 {
					directories = append(directories, arg)
				}
			}
		} else {
			fmt.Println("A valid directory path is required")
			return
		}

		verboseMode, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println("Failed to get value for verbose mode: " + err.Error())
			return
		}

		if _, err := ScanDirectories(directories, verboseMode); err != nil {
			fmt.Println("Directory Scan Failed. err:" + err.Error())
		}
	},
}

func ScanDirectories(directories []string, verboseMode bool) ([]*MatchDetails, error) {

	overAllMatchDetails := []*MatchDetails{}
	directoryMap, err := GetScanFileInfos(directories)
	if err != nil {
		return nil, fmt.Errorf("Failed to scan directories: " + err.Error())
	}
	totalDirectories := len(directoryMap)
	if totalDirectories == 0 {
		return nil, fmt.Errorf("no valid directories / files found for scanning")
	}

	dirNo := 0
	for _, fileMap := range directoryMap {
		dirNo++
		fileNo := 0
		totalFiles := len(fileMap)
		for filePath, fileInfo := range fileMap {
			fileNo++

			fmt.Printf("\n\n\n---\n\n\n")

			if verboseMode {
				fmt.Printf("\nScanning: Dir(%d of %d) File(%d of %d): %s\n", dirNo, totalDirectories, fileNo, totalFiles, filePath)
			}

			startTime := time.Now()
			matchDetails, fileLines, err := ScanFile(filePath)
			if err != nil {
				fmt.Printf("Failed to process file:%s err:%s\n", filePath, err.Error())
			}
			elapsed := time.Since(startTime)

			overAllMatchDetails = append(overAllMatchDetails, matchDetails...)
			for _, matchedDetails := range matchDetails {
				if prettyJson, err := json.MarshalIndent(*matchedDetails, "", "    "); err != nil {
					fmt.Printf("Failed to jsonify object: %+v err: %s\n", *matchedDetails, err.Error())
				} else {
					fmt.Printf("%s\n", string(prettyJson))
				}
			}

			if verboseMode {
				fileSize := fileInfo.Size()
				sizeStr := fmt.Sprintf("%d Byte(s)", fileSize)
				if fileSize > 1024*1024 {
					sizeStr = fmt.Sprintf("%d MB(s)", fileSize/(1024*1024))
				} else if fileSize > 1024 {
					sizeStr = fmt.Sprintf("%d KB(s)", fileSize/1024)
				}
				fmt.Printf("\nScan Complete: Dir(%d of %d) File(%d of %d): %s, size: %s, line(s): %d, timeTaken: %d ms\n",
					dirNo, totalDirectories, fileNo, totalFiles, filePath, sizeStr, fileLines, elapsed.Milliseconds())
			}
		}
	}

	return overAllMatchDetails, nil
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolP("verbose", "v", false, "")
}

func GetScanFileInfos(directories []string) (map[string]map[string]fs.FileInfo, error) {
	directoryMap := map[string]map[string]fs.FileInfo{}
	for _, directory := range directories {
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			fmt.Printf("No File / Directory found at %s. Skipping...", directory)
			continue
		}

		directoryMap[directory] = map[string]fs.FileInfo{}

		filepath.Walk(directory, func(filePath string, info fs.FileInfo, err error) error {

			if !info.IsDir() {
				directoryMap[directory][filePath] = info
			}
			return nil
		})
	}

	return directoryMap, nil
}

func ScanFile(scanFilePath string) ([]*MatchDetails, int, error) {

	scanFile, err := os.Open(scanFilePath)
	if err != nil {
		return nil, 0, fmt.Errorf("Failed to read file: " + scanFilePath + " err: " + err.Error())
	}
	defer scanFile.Close()

	matchDetailsList := []*MatchDetails{}

	scanner := bufio.NewScanner(scanFile)
	base58Re := regexp.MustCompile(`[a-km-zA-HJ-NP-Z1-9]+`)
	const maxCapacity = 256 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	linesProcessed := 0
	for scanner.Scan() {
		linesProcessed++
		fileLine := scanner.Text()
		if len(fileLine) > 0 {
			for _, word := range strings.Split(fileLine, " ") {
				matchedString := findLongestMatch(base58Re, word)
				matchedStringLength := len(matchedString)
				if len(word) > minPerfectMatchLength && matchedStringLength > minPossibleMatchLength {
					if matchedStringLength == len(word) {
						matchDetailsList = append(matchDetailsList, &MatchDetails{MatchType: "perfect", MatchFile: scanFilePath, MatchedLine: fileLine, MatchedLineNo: linesProcessed, MatchedWord: word})
					} else {
						matchDetailsList = append(matchDetailsList, &MatchDetails{MatchType: "possible", MatchFile: scanFilePath, MatchedLine: fileLine, MatchedLineNo: linesProcessed, MatchedWord: word})
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Occurred while scanning file: " + scanFilePath + "err: " + err.Error())
		return nil, linesProcessed, err
	}

	return matchDetailsList, linesProcessed, nil
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
