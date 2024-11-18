.PHONY: test coverage bench

test:
	go test -v ./pkg/...

coverage:
	go test -coverprofile=coverage.out ./pkg/...
	go tool cover -html=coverage.out

bench:
	go test -bench=. ./pkg/...