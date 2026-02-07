package rules

import (
	"regexp"
)

type HerokuNotCommandRule struct{}

func (r *HerokuNotCommandRule) ID() string { return "heroku_not_command" }

func (r *HerokuNotCommandRule) Match(command string, output string) bool {
	return regexp.MustCompile(`Run heroku _ to run`).MatchString(output)
}

func (r *HerokuNotCommandRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Run heroku _ to run ([^.]*)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	return command
}
