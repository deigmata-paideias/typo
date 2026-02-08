package common

type SlLsRule struct{}

func (r *SlLsRule) ID() string {
	return "sl_ls"
}

func (r *SlLsRule) Match(command string, output string) bool {
	return command == "sl"
}

func (r *SlLsRule) GetNewCommand(command string, output string) string {
	return "ls"
}
