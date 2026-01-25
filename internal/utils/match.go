package utils

import (
	"sort"

	"github.com/antlabs/strsim"
	"github.com/deigmata-paideias/typo/internal/types"
)

// Match returns the best matching command
func Match(str string, commands []string) string {
	bestMatchOne := strsim.FindBestMatch(str, commands)
	return commands[bestMatchOne.BestIndex]
}

// MatchMultiple returns multiple similar commands, sorted by similarity
func MatchMultiple(str string, commands []string, limit int) []types.MatchResult {
	if limit <= 0 {
		limit = 5
	}

	matches := strsim.FindBestMatch(str, commands)
	results := make([]types.MatchResult, 0, len(matches.AllResult))

	for _, rating := range matches.AllResult {
		results = append(results, types.MatchResult{
			Command: rating.S,
			Score:   rating.Score,
		})
	}

	// Sort by score in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	// Limit the number of results returned
	if len(results) > limit {
		results = results[:limit]
	}

	return results
}

// SortAndLimitResults sorts results by score and limits the number of results
func SortAndLimitResults(results []types.MatchResult, limit int) []types.MatchResult {
	if limit <= 0 {
		limit = 5
	}

	// Sort by score in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	// Remove duplicates based on command
	seen := make(map[string]bool)
	uniqueResults := make([]types.MatchResult, 0, len(results))
	for _, result := range results {
		if !seen[result.Command] {
			seen[result.Command] = true
			uniqueResults = append(uniqueResults, result)
		}
	}

	// Limit the number of results
	if len(uniqueResults) > limit {
		uniqueResults = uniqueResults[:limit]
	}

	return uniqueResults
}
