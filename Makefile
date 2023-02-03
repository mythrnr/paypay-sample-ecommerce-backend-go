ifndef VERBOSE
MAKEFLAGS += --silent
endif

.PHONY: clean
clean:
	rm -rf .cache/*

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	docker pull golangci/golangci-lint:latest > /dev/null \
	&& mkdir -p .cache/golangci-lint \
	&& docker run --rm \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/root/.cache \
		-w /app golangci/golangci-lint:latest golangci-lint run ./...

.PHONY: nancy
nancy:
	docker pull sonatypecommunity/nancy:latest > /dev/null \
	&& go list -buildvcs=false -deps -json ./... \
	| docker run --rm -i sonatypecommunity/nancy:latest sleuth

.PHONY: refund
refund:
	go run cmd/refund/main.go

.PHONY: serve
serve:
	go run shop/main.go

.PHONY: spell-check
spell-check:
	docker pull ghcr.io/streetsidesoftware/cspell:latest > /dev/null \
	&& docker run --rm \
		-v $(shell pwd):/workdir \
		ghcr.io/streetsidesoftware/cspell:latest \
			--config .vscode/cspell.json "**"

.PHONY: tidy
tidy:
	go mod tidy
