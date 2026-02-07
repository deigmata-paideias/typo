package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type PortAlreadyInUseRule struct{}

func (r *PortAlreadyInUseRule) ID() string { return "port_already_in_use" }

func (r *PortAlreadyInUseRule) Match(command string, output string) bool {
	patterns := []string{
		`bind on address \('.*', (\d+)\)`,
		`Unable to bind [^ ]*:(\d+)`,
		`can't listen on port (\d+)`,
		`listen EADDRINUSE [^ ]*:(\d+)`,
		`Address already in use`, // Generic
	}
	for _, p := range patterns {
		if regexp.MustCompile(p).MatchString(output) {
			return true
		}
	}
	return false
}

func (r *PortAlreadyInUseRule) GetNewCommand(command string, output string) string {
	patterns := []string{
		`bind on address \('.*', (\d+)\)`,
		`Unable to bind [^ ]*:(\d+)`,
		`can't listen on port (\d+)`,
		`listen EADDRINUSE [^ ]*:(\d+)`,
	}
	var port string
	for _, p := range patterns {
		re := regexp.MustCompile(p)
		matches := re.FindStringSubmatch(output)
		if len(matches) > 1 {
			port = matches[1]
			break
		}
	}

	if port == "" {
		return command
	}

	// Find PID using lsof
	// lsof -i :<port>
	// Output:
	// COMMAND PID USER ...
	// node 1234 user ...
	out, err := utils.ExecCommandWithOutput("lsof", "-i", ":"+port)
	if err != nil {
		return command
	}

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) > 1 {
		fields := strings.Fields(lines[1])
		if len(fields) > 1 {
			pid := fields[1]
			return fmt.Sprintf("kill %s && %s", pid, command)
		}
	}
	return command
}
