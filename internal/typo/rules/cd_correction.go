package rules

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type CdCorrectionRule struct{}

func (r *CdCorrectionRule) ID() string {
	return "cd_correction"
}

func (r *CdCorrectionRule) Match(command string, output string) bool {
	if !strings.HasPrefix(command, "cd ") {
		return false
	}
	lowerOutput := strings.ToLower(output)
	return strings.Contains(lowerOutput, "no such file or directory") ||
		strings.Contains(lowerOutput, "test") || // bash/zsh specific?
		strings.Contains(lowerOutput, "does not exist") ||
		strings.Contains(lowerOutput, "can't cd to")
}

func (r *CdCorrectionRule) GetNewCommand(command string, output string) string {
	// Attempt to correct path
	// cd /foo/bar/baz
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}
	pathArg := parts[1]

	// Logic: walk path components
	// Absolute or relative?
	cwd, _ := os.Getwd()

	// Split path
	// Handle separators
	pathParts := strings.Split(pathArg, string(os.PathSeparator))

	// Start root
	currentSearchPath := cwd
	if strings.HasPrefix(pathArg, string(os.PathSeparator)) {
		currentSearchPath = string(os.PathSeparator)
		// If first part is empty string due to split "/foo" -> ["", "foo"]
		if len(pathParts) > 0 && pathParts[0] == "" {
			pathParts = pathParts[1:]
		}
	} else {
		// Relative
	}

	// Reconstruct
	var correctedParts []string
	if currentSearchPath == string(os.PathSeparator) {
		correctedParts = append(correctedParts, "") // for join later
	}

	// NOTE: This logic is tricky to port 1:1 reliably without more extensive testing.
	// The python version handles it component by component.
	// I will implement a simplified version: try to correct the basename if dirname exists.

	// dir := filepath.Dir(pathArg)
	// base := filepath.Base(pathArg)

	// Check if dir exists
	// If path is "foo/bar", dir is "foo".
	// If path is "bar", dir is ".".

	// Simplified logic: If the parent exists, try to match the child.
	// Full path walking is safer but complex.

	// Let's defer to CdMkdirRule logic which suggests "mkdir -p".
	// But this rule is about TYPOS.

	// Let's implement the iterative walk properly.

	searchRoot := cwd
	if filepath.IsAbs(pathArg) {
		searchRoot = "/" // Unix assumption
	}

	// Split again cleanly
	cleanPath := filepath.Clean(pathArg)
	components := strings.Split(cleanPath, string(os.PathSeparator))

	var finalPath []string
	if filepath.IsAbs(pathArg) {
		// Start from root
		searchRoot = "/"
		// components of "/a/b" are ["", "a", "b"] ?
		// filepath.Clean("/a/b") -> "/a/b"
	} else {
		searchRoot = cwd
	}

	// Iterate components
	curr := searchRoot
	for _, comp := range components {
		if comp == "" || comp == "." {
			continue
		}
		if comp == ".." {
			curr = filepath.Dir(curr)
			finalPath = append(finalPath, "..")
			continue
		}

		// Check if comp exists in curr
		full := filepath.Join(curr, comp)
		if _, err := os.Stat(full); err == nil {
			// Exists
			curr = full
			finalPath = append(finalPath, comp)
		} else {
			// Mistyped?
			entries, err := os.ReadDir(curr)
			if err != nil {
				return command // permission or other error
			}

			var candidates []string
			for _, e := range entries {
				if e.IsDir() {
					candidates = append(candidates, e.Name())
				}
			}

			best := utils.Match(comp, candidates)
			if best != "" {
				curr = filepath.Join(curr, best)
				finalPath = append(finalPath, best)
			} else {
				// No match, maybe fallback to cd_mkdir logic?
				// The python rule falls back to cd_mkdir.
				return command // Or implement cd_mkdir fallback logic here?
			}
		}
	}

	// Reassemble
	result := filepath.Join(finalPath...)
	if filepath.IsAbs(pathArg) {
		result = "/" + result
	}

	return "cd " + result
}
