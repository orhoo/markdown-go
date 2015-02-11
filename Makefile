.PHONY: build


build:
	go install ./...

test:
	go test -race ./...

