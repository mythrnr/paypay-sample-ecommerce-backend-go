.PHONY: serve tidy

refund:
	go run cmd/refund/main.go

serve:
	go run shop/main.go

tidy:
	go mod tidy
