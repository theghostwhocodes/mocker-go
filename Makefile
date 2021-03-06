.PHONY: build test cover-report

default: build

build:
	go build -o build/mocker cmd/mocker/main.go

test: test_unit test_e2e

test_unit:
	go test github.com/theghostwhocodes/mocker-go/internal/... -coverprofile=cover.out

test_content_managers:
	go test github.com/theghostwhocodes/mocker-go/internal/contentManagers -v -coverprofile=cover.out

test_filters:
	go test github.com/theghostwhocodes/mocker-go/internal/filters -v -coverprofile=cover.out

test_validators:
	go test github.com/theghostwhocodes/mocker-go/internal/validators -v -coverprofile=cover.out

test_e2e:
	go test github.com/theghostwhocodes/mocker-go/tests/e2e/... -coverprofile=cover.out

cover-report:
	go tool cover -html=cover.out
