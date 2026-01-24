package utils

import (
	"testing"
)

// not normal
func Test_Match(t *testing.T) {

	// skip 
	t.Skip()

	var commands = []string{"pkgutil", "pkgutils", "ab", "git", "asfg", "vim", "branch"}
	data := []struct {
		Name     string
		Str      string
		Expected string
	}{
		{"git-1", "gti", "git"},
		{"git-2", "tgi", "git"},
		{"git-3", "tig", "git"},
	}

	for _, datum := range data {
		if Match(datum.Str, commands) != datum.Expected {
			t.Errorf("Test case: %v, Expected %s but got %s", datum.Name, datum.Expected, Match(datum.Str, commands))
		}
	}

}
