SHELL:=/bin/bash

DATETIME = $(shell date +"%Y%m%d%H%M%S")

# Log the running target
LOG_TARGET = echo -e "\033[0;32m==================> Running $@ ============> ... \033[0m"

# Log debugging info
define log
echo -e "\033[36m==================>$1\033[0m"
endef

# Log error info
define errorLog
echo -e "\033[0;31m==================>$1\033[0m"
endef

.PHONY: help
help:
	@echo -e "\033[1;3;34m Like thefuck, but he uses Go to implement it more intelligently.\033[0m\n"
	@echo -e "Usage:\n  make \033[36m<Target>\033[0m \033[36m<Option>\033[0m\n\nTargets:"
	@awk 'BEGIN {FS = ":.*##"; printf ""} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
