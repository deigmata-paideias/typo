package rules

// Stub for ChocoInstallRule to prevent build error if removed
type ChocoInstallRule struct{}

func (r *ChocoInstallRule) ID() string                                         { return "choco_install" }
func (r *ChocoInstallRule) Match(command string, output string) bool           { return false }
func (r *ChocoInstallRule) GetNewCommand(command string, output string) string { return command }

// Stub for DirtyUntarRule
type DirtyUntarRule struct{}

func (r *DirtyUntarRule) ID() string                                         { return "dirty_untar" }
func (r *DirtyUntarRule) Match(command string, output string) bool           { return false }
func (r *DirtyUntarRule) GetNewCommand(command string, output string) string { return command }

// Stub for FixFileRule
type FixFileRule struct{}

func (r *FixFileRule) ID() string                                         { return "fix_file" }
func (r *FixFileRule) Match(command string, output string) bool           { return false }
func (r *FixFileRule) GetNewCommand(command string, output string) string { return command }

// Stubs for Git Rules Recovery
type GemUnknownCommandRule struct{}

func (r *GemUnknownCommandRule) ID() string                                         { return "gem_unknown_command" }
func (r *GemUnknownCommandRule) Match(command string, output string) bool           { return false }
func (r *GemUnknownCommandRule) GetNewCommand(command string, output string) string { return command }

type GitBranch0FlagRule struct{}

func (r *GitBranch0FlagRule) ID() string                                         { return "git_branch_0flag" }
func (r *GitBranch0FlagRule) Match(command string, output string) bool           { return false }
func (r *GitBranch0FlagRule) GetNewCommand(command string, output string) string { return command }

type GitCloneGitCloneRule struct{}

func (r *GitCloneGitCloneRule) ID() string                                         { return "git_clone_git_clone" }
func (r *GitCloneGitCloneRule) Match(command string, output string) bool           { return false }
func (r *GitCloneGitCloneRule) GetNewCommand(command string, output string) string { return command }

type GitCommitAmendRule struct{}

func (r *GitCommitAmendRule) ID() string                                         { return "git_commit_amend" }
func (r *GitCommitAmendRule) Match(command string, output string) bool           { return false }
func (r *GitCommitAmendRule) GetNewCommand(command string, output string) string { return command }

type GitDiffNoIndexRule struct{}

func (r *GitDiffNoIndexRule) ID() string                                         { return "git_diff_no_index" }
func (r *GitDiffNoIndexRule) Match(command string, output string) bool           { return false }
func (r *GitDiffNoIndexRule) GetNewCommand(command string, output string) string { return command }

type GitFixStashRule struct{}

func (r *GitFixStashRule) ID() string                                         { return "git_fix_stash" }
func (r *GitFixStashRule) Match(command string, output string) bool           { return false }
func (r *GitFixStashRule) GetNewCommand(command string, output string) string { return command }

type GitMainMasterRule struct{}

func (r *GitMainMasterRule) ID() string                                         { return "git_main_master" }
func (r *GitMainMasterRule) Match(command string, output string) bool           { return false }
func (r *GitMainMasterRule) GetNewCommand(command string, output string) string { return command }

type GitMergeRule struct{}

func (r *GitMergeRule) ID() string                                         { return "git_merge" }
func (r *GitMergeRule) Match(command string, output string) bool           { return false }
func (r *GitMergeRule) GetNewCommand(command string, output string) string { return command }

type GitPullRule struct{}

func (r *GitPullRule) ID() string                                         { return "git_pull" }
func (r *GitPullRule) Match(command string, output string) bool           { return false }
func (r *GitPullRule) GetNewCommand(command string, output string) string { return command }

type GitPushPullRule struct{}

func (r *GitPushPullRule) ID() string                                         { return "git_push_pull" }
func (r *GitPushPullRule) Match(command string, output string) bool           { return false }
func (r *GitPushPullRule) GetNewCommand(command string, output string) string { return command }

// Batch 3 Stubs
type GitPushWithoutCommitsRule struct{}

func (r *GitPushWithoutCommitsRule) ID() string                               { return "git_push_without_commits" }
func (r *GitPushWithoutCommitsRule) Match(command string, output string) bool { return false }
func (r *GitPushWithoutCommitsRule) GetNewCommand(command string, output string) string {
	return command
}

type GitRmLocalModificationsRule struct{}

func (r *GitRmLocalModificationsRule) ID() string                               { return "git_rm_local_modifications" }
func (r *GitRmLocalModificationsRule) Match(command string, output string) bool { return false }
func (r *GitRmLocalModificationsRule) GetNewCommand(command string, output string) string {
	return command
}

type GulpNotTaskRule struct{}

func (r *GulpNotTaskRule) ID() string                                         { return "gulp_not_task" }
func (r *GulpNotTaskRule) Match(command string, output string) bool           { return false }
func (r *GulpNotTaskRule) GetNewCommand(command string, output string) string { return command }

type HasExistsScriptRule struct{}

func (r *HasExistsScriptRule) ID() string                                         { return "has_exists_script" }
func (r *HasExistsScriptRule) Match(command string, output string) bool           { return false }
func (r *HasExistsScriptRule) GetNewCommand(command string, output string) string { return command }

type JavacRule struct{}

func (r *JavacRule) ID() string                                         { return "javac" }
func (r *JavacRule) Match(command string, output string) bool           { return false }
func (r *JavacRule) GetNewCommand(command string, output string) string { return command }

type LeinNotTaskRule struct{}

func (r *LeinNotTaskRule) ID() string                                         { return "lein_not_task" }
func (r *LeinNotTaskRule) Match(command string, output string) bool           { return false }
func (r *LeinNotTaskRule) GetNewCommand(command string, output string) string { return command }

type ManNoSpaceRule struct{}

func (r *ManNoSpaceRule) ID() string                                         { return "man_no_space" }
func (r *ManNoSpaceRule) Match(command string, output string) bool           { return false }
func (r *ManNoSpaceRule) GetNewCommand(command string, output string) string { return command }

type NpmMissingScriptRule struct{}

func (r *NpmMissingScriptRule) ID() string                                         { return "npm_missing_script" }
func (r *NpmMissingScriptRule) Match(command string, output string) bool           { return false }
func (r *NpmMissingScriptRule) GetNewCommand(command string, output string) string { return command }

type NpmRunScriptRule struct{}

func (r *NpmRunScriptRule) ID() string                                         { return "npm_run_script" }
func (r *NpmRunScriptRule) Match(command string, output string) bool           { return false }
func (r *NpmRunScriptRule) GetNewCommand(command string, output string) string { return command }

type OmnienvNoSuchCommandRule struct{}

func (r *OmnienvNoSuchCommandRule) ID() string                               { return "omnienv_no_such_command" }
func (r *OmnienvNoSuchCommandRule) Match(command string, output string) bool { return false }
func (r *OmnienvNoSuchCommandRule) GetNewCommand(command string, output string) string {
	return command
}

// Batch 4 Stubs
type OpenRule struct{}

func (r *OpenRule) ID() string                                         { return "open" }
func (r *OpenRule) Match(command string, output string) bool           { return false }
func (r *OpenRule) GetNewCommand(command string, output string) string { return command }

type PacmanInvalidOptionRule struct{}

func (r *PacmanInvalidOptionRule) ID() string                                         { return "pacman_invalid_option" }
func (r *PacmanInvalidOptionRule) Match(command string, output string) bool           { return false }
func (r *PacmanInvalidOptionRule) GetNewCommand(command string, output string) string { return command }

type PythonCommandRule struct{}

func (r *PythonCommandRule) ID() string                                         { return "python_command" }
func (r *PythonCommandRule) Match(command string, output string) bool           { return false }
func (r *PythonCommandRule) GetNewCommand(command string, output string) string { return command }
