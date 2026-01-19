##@ Linter

.PHONY: lint
lint: ## Check files
lint: markdown-lint-check yaml-lint codespell

.PHONY: codespell
codespell: CODESPELL_SKIP := $(shell cat tools/linter/codespell/.codespell.skip | tr \\n ',')
codespell: ## Check the code-spell
	@$(LOG_TARGET)
	codespell --version
	codespell --skip "$(CODESPELL_SKIP)" --ignore-words ./tools/linter/codespell/.codespell.ignorewords

.PHONY: yaml-lint
yaml-lint: ## Check the yaml lint
	@$(LOG_TARGET)
	yamllint --version
	yamllint -c ./tools/linter/yamllint/.yamllint .

.PHONY: yaml-lint-fix
yaml-lint-fix: ## Yaml lint fix
	@$(LOG_TARGET)
	yamlfmt -version
	yamlfmt .

.PHONY: markdown-lint-check
markdown-lint-check: ## Check the markdown files.
	@$(LOG_TARGET)
	markdownlint --version
	markdownlint --config ./tools/linter/markdownlint/markdown_lint_config.yaml .

.PHONY: markdown-lint-fix
markdown-lint-fix: ## Fix the markdown files style.
	@$(LOG_TARGET)
	markdownlint --version
	markdownlint --config ./tools/linter/markdownlint/markdown_lint_config.yaml --fix .
