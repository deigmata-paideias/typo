package common

import (
	"strings"
)

type QuotationMarksRule struct{}

func (r *QuotationMarksRule) ID() string { return "quotation_marks" }

func (r *QuotationMarksRule) Match(command string, output string) bool {
	return strings.Contains(command, "'") && strings.Contains(command, "\"")
}

func (r *QuotationMarksRule) GetNewCommand(command string, output string) string {
	return strings.ReplaceAll(command, "'", "\"")
}
