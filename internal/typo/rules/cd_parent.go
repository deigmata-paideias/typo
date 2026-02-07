package rules

type CdParentRule struct{}

func (r *CdParentRule) ID() string {
	return "cd_parent"
}

func (r *CdParentRule) Match(command string, output string) bool {
	return command == "cd.."
}

func (r *CdParentRule) GetNewCommand(command string, output string) string {
	return "cd .."
}
