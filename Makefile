.PHONY: build test

default: build

build:
	go build -o build/mocker cmd/mocker/main.go

test:
	go test github.com/theghostwhocodes/mocker-go/internal/... -coverprofile=cover.out

test_content_managers:
	go test github.com/theghostwhocodes/mocker-go/internal/contentManagers -coverprofile=cover.out

test_validators:
	go test github.com/theghostwhocodes/mocker-go/internal/validators -coverprofile=cover.out