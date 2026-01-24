package typo

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/types"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type ITypo interface {
	// Typo 返回修正的命令候选列表和原始命令
	Typo() (string, []types.MatchResult, error)
}

type LocalTypo struct {
	repo repository.IRepository
	hs   scanner.IScanner
}

func NewLocalTypo(repo repository.IRepository, hs scanner.IScanner) ITypo {

	return &LocalTypo{
		repo,
		hs,
	}
}

// Typo 返回修正的命令候选列表和原始命令
// first return: original command
// second return: candidate commands with scores
func (t *LocalTypo) Typo() (string, []types.MatchResult, error) {

	command, err := t.hs.Scan()
	if err != nil {
		return "", nil, err
	}

	// 获取命令的第一个词
	oneCmd := strings.Split(command, " ")[0]

	commandNames, err := t.repo.GetAllCommandNames()
	if err != nil {
		return "", nil, err
	}

	matches := utils.MatchMultiple(oneCmd, commandNames, 5)

	if len(matches) > 0 && matches[0].Score < 0.5 {
		_, err := t.repo.FindCommandByName(oneCmd)
		if err == nil {
			return command, nil, nil
		}
	}

	// add command description
	for i := range matches {
		cmd, err := t.repo.FindCommandByName(matches[i].Command)
		if err != nil {
			matches[i].Desc = ""
			continue
		}
		matches[i].Desc = cmd.Description
	}

	return command, matches, nil
}
