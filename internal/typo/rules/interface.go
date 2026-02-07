package rules

// Rule interface defines a correction rule
type Rule interface {
	// ID returns the unique identifier of the rule
	ID() string
	// Match checks if the rule applies to the given command and its output
	Match(command string, output string) bool
	// GetNewCommand returns the corrected command
	GetNewCommand(command string, output string) string
}
