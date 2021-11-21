package main

import (
	"fmt"
	"testing"

	"github.com/adnan-mansoor-2015/scan-bitcoin-addresses/cmd"
	"github.com/stretchr/testify/require"
)

// Testing Possible Matches
func TestPerfectMatches(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/perfect")
	require.NoError(t, err, "Failed to scan directory")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "perfect" {
			require.Equal(t, "perfect", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Perfect Matches
func TestPossibleMatches(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/possible")
	require.NoError(t, err, "Failed to scan directory")
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchType != "possible" {
			require.Equal(t, "possible", matchDetail.MatchType, "UnExpected Match: %+v", matchDetail)
		}
	}
}

// Testing Possible Matches
func TestNoMatchesMatches(t *testing.T) {
	matchDetails, err := cmd.ScanDirectory("testDirectory/nomatch")
	require.NoError(t, err, "Failed to scan directory")
	require.Equal(t, 0, len(matchDetails), "Matches found somehow")

	for _, matchDetail := range matchDetails {
		fmt.Printf("%+v\n", matchDetail)
	}
}
