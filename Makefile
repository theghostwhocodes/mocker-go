.PHONY: build test

default: build

build:
	go build -o build/mocker cmd/mocker/main.go

test:
	go test github.com/theghostwhocodes/mocker-go/internal/... -coverprofile=cover.out
