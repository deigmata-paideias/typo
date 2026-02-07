package rules

import (
	"regexp"
	"strings"
)

type RailsMigrationsPendingRule struct{}

func (r *RailsMigrationsPendingRule) ID() string { return "rails_migrations_pending" }

func (r *RailsMigrationsPendingRule) Match(command string, output string) bool {
	return strings.Contains(output, "Migrations are pending. To resolve this issue, run:")
}

func (r *RailsMigrationsPendingRule) GetNewCommand(command string, output string) string {
	// To resolve this issue, run:
	// db:migrate rais...
	re := regexp.MustCompile(`To resolve this issue, run:\s+(.*?)\n`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1] + " && " + command
	}
	return command
}
