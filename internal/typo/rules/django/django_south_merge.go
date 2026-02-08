package django

import (
	"strings"
)

type DjangoSouthMergeRule struct{}

func (r *DjangoSouthMergeRule) ID() string {
	return "django_south_merge"
}

func (r *DjangoSouthMergeRule) Match(command string, output string) bool {
	return strings.Contains(command, "manage.py") &&
		strings.Contains(command, "migrate") &&
		strings.Contains(output, "--merge: will just attempt the migration")
}

func (r *DjangoSouthMergeRule) GetNewCommand(command string, output string) string {
	return command + " --merge"
}
