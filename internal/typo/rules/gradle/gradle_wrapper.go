package gradle

import (
	"os"
	"strings"
)

type GradleWrapperRule struct{}

func (r *GradleWrapperRule) ID() string { return "gradle_wrapper" }

func (r *GradleWrapperRule) Match(command string, output string) bool {
	if !strings.HasPrefix(command, "gradle") {
		return false
	}
	// Check if ./gradlew exists
	if _, err := os.Stat("./gradlew"); os.IsNotExist(err) {
		return false
	}
	// And 'not found' in output (which implies gradle command itself wasn't found or failed)
	// The Python rule uses `not which(...)` AND `'not found' in output`.
	// For simplicity, if gradle command failed and gradlew exists, suggest it.
	return strings.Contains(output, "not found") || strings.Contains(output, "No such file or directory")
}

func (r *GradleWrapperRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "gradle", "./gradlew", 1)
}
