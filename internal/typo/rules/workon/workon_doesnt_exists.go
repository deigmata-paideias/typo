package workon

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type WorkonDoesntExistsRule struct{}

func (r *WorkonDoesntExistsRule) ID() string { return "workon_doesnt_exists" }

func (r *WorkonDoesntExistsRule) Match(command string, output string) bool {
	// workon usage: workon <env>
	// if <env> does not exist.
	// We check if command starts with workon and arg count.
	parts := strings.Fields(command)
	if len(parts) < 2 || parts[0] != "workon" {
		return false
	}
	env := parts[1]

	// Check if env exists in ~/.virtualenvs
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	venvRoot := filepath.Join(home, ".virtualenvs")
	if _, err := os.Stat(filepath.Join(venvRoot, env)); err != nil {
		// Env doesn't exist
		return true
	}
	return false
}

func (r *WorkonDoesntExistsRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	wrongEnv := parts[1]

	// Suggest creating it?
	// return "mkvirtualenv " + wrongEnv

	// Or suggest similar existing envs
	home, _ := os.UserHomeDir()
	venvRoot := filepath.Join(home, ".virtualenvs")
	files, err := os.ReadDir(venvRoot)
	var candidates []string
	if err == nil {
		for _, f := range files {
			if f.IsDir() {
				candidates = append(candidates, f.Name())
			}
		}
	}

	if len(candidates) > 0 {
		best := utils.Match(wrongEnv, candidates)
		if best != "" {
			return strings.Replace(command, wrongEnv, best, 1) // First Candidate
		}
	}

	// Fallback to create
	return "mkvirtualenv " + wrongEnv
}
