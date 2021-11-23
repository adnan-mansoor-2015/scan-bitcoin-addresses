package main

import (
	"fmt"
	"testing"

	"github.com/adnan-mansoor-2015/scan-bitcoin-addresses/cmd"
	"github.com/stretchr/testify/require"
)

// Testing Possible Matches Directory
func TestPerfectMatchesDirectory(t *testing.T) {
	matchDetails, err := cmd.ScanDirectories([]string{"testDirectory/perfect"}, true)
	require.NoError(t, err, "Failed to scan directory")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "perfect" {
			require.Equal(t, "perfect", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Perfect Matches Directory
func TestPossibleMatchesDirectory(t *testing.T) {
	matchDetails, err := cmd.ScanDirectories([]string{"testDirectory/possible"}, true)
	require.NoError(t, err, "Failed to scan directory")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "possible" {
			require.Equal(t, "possible", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Possible Matches Directory
func TestNoMatchesMatchesDirectory(t *testing.T) {
	matchDetails, err := cmd.ScanDirectories([]string{"testDirectory/nomatch"}, true)
	require.NoError(t, err, "Failed to scan directory")
	require.Equal(t, 0, len(matchDetails), "Matches found somehow")

	for _, matchDetail := range matchDetails {
		fmt.Printf("Unexpected Match: %+v\n", matchDetail)
	}
}

// Testing Possible Matches File
func TestPerfectMatchesFile(t *testing.T) {
	matchDetails, err := cmd.ScanDirectories([]string{"testDirectory/perfect/btc/perfect.txt"}, true)
	require.NoError(t, err, "Failed to scan file")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "perfect" {
			require.Equal(t, "perfect", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Perfect Matches File
func TestPossibleMatchesFile(t *testing.T) {
	matchDetails, err := cmd.ScanDirectories([]string{"testDirectory/possible/btc/possible.txt"}, true)
	require.NoError(t, err, "Failed to scan file")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "possible" {
			require.Equal(t, "possible", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Possible Matches File
func TestNoMatchesMatchesFile(t *testing.T) {
	matchDetails, err := cmd.ScanDirectories([]string{"testDirectory/nomatch/btc/nomatch.txt"}, true)
	require.NoError(t, err, "Failed to scan file")
	require.Equal(t, 0, len(matchDetails), "Matches found somehow")

	for _, matchDetail := range matchDetails {
		fmt.Printf("Unexpected Match: %+v\n", matchDetail)
	}
}
