package utils

import (
	"sort"

	"github.com/antlabs/strsim"
	"github.com/deigmata-paideias/typo/internal/types"
)

// Match 返回最佳匹配命令
func Match(str string, commands []string) string {
	bestMatchOne := strsim.FindBestMatch(str, commands)
	return commands[bestMatchOne.BestIndex]
}

// MatchMultiple 返回多个相似命令，按相似度排序
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

	// 按分数降序排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	// 限制返回数量
	if len(results) > limit {
		results = results[:limit]
	}

	return results
}
