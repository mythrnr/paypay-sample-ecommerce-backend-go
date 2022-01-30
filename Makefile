.PHONY: refund
.SILENT: refund
refund:
	go run cmd/refund/main.go

.PHONY: serve
.SILENT: serve
serve:
	go run shop/main.go

.PHONY: tidy
.SILENT: tidy
tidy:
	go mod tidy
