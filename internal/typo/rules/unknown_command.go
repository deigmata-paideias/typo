package rules

import (
	"regexp"
	"strings"
)

type UnknownCommandRule struct{}

func (r *UnknownCommandRule) ID() string { return "unknown_command" }

func (r *UnknownCommandRule) Match(command string, output string) bool {
	// Pattern: <prog>: Unknown command ... Did you mean ...
	return strings.Contains(output, ": Unknown command") && strings.Contains(output, "Did you mean")
}

func (r *UnknownCommandRule) GetNewCommand(command string, output string) string {
	reBroken := regexp.MustCompile(`([^:]*): Unknown command`)
	reFix := regexp.MustCompile(`Did you mean ([^?]*)`)

	brokenMat := reBroken.FindStringSubmatch(output)
	fixMat := reFix.FindStringSubmatch(output)

	if len(brokenMat) > 1 && len(fixMat) > 1 {
		// brokenMat[1] matches "git", "foo", implies PROGRAM name?
		// No, usually "program: Unknown command 'subcommand'"
		// Wait, the regex `([^:]*): Unknown command` captures the PROGRAM name if the error is "git: Unknown command...".
		// But usually we want to correct the SUBCOMMAND.
		// The regex in python is: `broken_cmd = re.findall(r"([^:]*): Unknown command.*", command.output)[0]`
		// If output is `git: Unknown command 'commt'`, match is `git`.
		// replace_command(command, broken_cmd, matched)
		// Wait, if it replaces `git` with `commit`? That's wrong.
		// `replace_command` in `thefuck` replaces the token that matches.
		// If command is `git commt`. `broken_cmd` is `git` (from regex).
		// That seems suspicious.
		// Let's re-read python: `broken_cmd = re.findall(r"([^:]*): Unknown command.*", command.output)[0]`
		// If output is `program: Unknown command "subcmd"`.
		// Python re `[^:]*` matches `program`.
		// Then it calls `replace_command(command, broken_cmd, matched)`.
		// This suggests it replaces `program`?
		// Ah, `unknown_command.py` is generic.
		// If I type `myprog flarg`. Output: `myprog: Unknown command flarg`.
		// If I type `giit status`. Output: `giit: Unknown command`.
		// Maybe it fixes the program name itself?
		// "Did you mean git?"
		// Yes, likely fixing the binary/program name.

		broken := strings.TrimSpace(brokenMat[1])
		fix := strings.TrimSpace(fixMat[1])
		return strings.Replace(command, broken, fix, 1)
	}
	return command
}
