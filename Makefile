ifndef VERBOSE
MAKEFLAGS += --silent
endif

.PHONY: refund
refund:
	go run cmd/refund/main.go

.PHONY: serve
serve:
	go run shop/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: nancy
nancy:
	go list -json -m all | nancy sleuth

.PHONY: spell-check
spell-check:
	# npm install -g cspell@latest
	cspell lint --config .vscode/cspell.json ".*" && \
	cspell lint --config .vscode/cspell.json "**/.*" && \
	cspell lint --config .vscode/cspell.json ".{github,vscode}/**/*" && \
	cspell lint --config .vscode/cspell.json "**"

.PHONY: tidy
tidy:
	go mod tidy
