.PHONY: setup
setup: install-static-tools
	cp .pre-commit .git/hooks/pre-commit

install-static-tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest