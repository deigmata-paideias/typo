##@ Tools

.PHONY: tools

tools: ## Install ci tools

	@$(LOG_TARGET)
	go version
	python --version
	node --version
	npm --version

	@echo "Setting up Python venv"
	pip install --upgrade pip

	@echo "Installing markdownlint-cli"
	npm install markdownlint-cli --global

	@echo "Installing golangci-lint"
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0

	@echo "Installing codespell"
	pip install codespell

	@echo "Installing yamllint"
	pip install yamllint==1.35.1

	@echo "Installing yamlfmt"
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
