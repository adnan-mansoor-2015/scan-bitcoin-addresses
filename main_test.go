package main

import (
	"fmt"
	"testing"

	"github.com/adnan-mansoor-2015/scan-bitcoin-addresses/cmd"
	"github.com/stretchr/testify/require"
)

// Testing Possible Matches Directory
func TestPerfectMatchesDirectory(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/perfect")
	require.NoError(t, err, "Failed to scan directory")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "perfect" {
			require.Equal(t, "perfect", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Perfect Matches Directory
func TestPossibleMatchesDirectory(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/possible")
	require.NoError(t, err, "Failed to scan directory")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "possible" {
			require.Equal(t, "possible", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Possible Matches Directory
func TestNoMatchesMatchesDirectory(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/nomatch")
	require.NoError(t, err, "Failed to scan directory")
	require.Equal(t, 0, len(matchDetails), "Matches found somehow")

	for _, matchDetail := range matchDetails {
		fmt.Printf("Unexpected Match: %+v\n", matchDetail)
	}
}

// Testing Possible Matches File
func TestPerfectMatchesFile(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/perfect/btc/perfect.txt")
	require.NoError(t, err, "Failed to scan file")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "perfect" {
			require.Equal(t, "perfect", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Perfect Matches File
func TestPossibleMatchesFile(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/possible/btc/possible.txt")
	require.NoError(t, err, "Failed to scan file")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "possible" {
			require.Equal(t, "possible", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Possible Matches File
func TestNoMatchesMatchesFile(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/nomatch/btc/nomatch.txt")
	require.NoError(t, err, "Failed to scan file")
	require.Equal(t, 0, len(matchDetails), "Matches found somehow")

	for _, matchDetail := range matchDetails {
		fmt.Printf("Unexpected Match: %+v\n", matchDetail)
	}
}
