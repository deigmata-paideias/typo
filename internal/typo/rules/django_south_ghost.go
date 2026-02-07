package rules

import (
	"strings"
)

type DjangoSouthGhostRule struct{}

func (r *DjangoSouthGhostRule) ID() string {
	return "django_south_ghost"
}

func (r *DjangoSouthGhostRule) Match(command string, output string) bool {
	return strings.Contains(command, "manage.py") &&
		strings.Contains(command, "migrate") &&
		strings.Contains(output, "or pass --delete-ghost-migrations")
}

func (r *DjangoSouthGhostRule) GetNewCommand(command string, output string) string {
	return command + " --delete-ghost-migrations"
}
