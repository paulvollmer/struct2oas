all: lint build

lint:
	@golint

build:
	@go fmt
	@go build

clean:
	rm -f *.yml

.PHONY: lint build clean

test: build test-files clean test-folder

test-files:
	./struct2oas -leftpad "  " -source fixtures/model1.go
	./struct2oas -leftpad "  " -source fixtures/model2.go
	./struct2oas -leftpad "  " -source fixtures/types_multi.go
	./struct2oas -leftpad "  " -source fixtures/types.go

test-folder:
	./struct2oas -leftpad "  " -source fixtures

.PHONY: test test-files test-folder

release:
	gox

.PHONY: release
