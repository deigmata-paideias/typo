package cargo

type CargoRule struct{}

func (r *CargoRule) ID() string {
	return "cargo"
}

func (r *CargoRule) Match(command string, output string) bool {
	return command == "cargo"
}

func (r *CargoRule) GetNewCommand(command string, output string) string {
	return "cargo build"
}
