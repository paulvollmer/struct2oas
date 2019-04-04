build:
	@go fmt
	@go build

test: build test-files clean test-folder

test-files:
	./struct2oas -leftpad "  " -source fixtures/model1.go
	./struct2oas -leftpad "  " -source fixtures/model2.go
	./struct2oas -leftpad "  " -source fixtures/types_multi.go
	./struct2oas -leftpad "  " -source fixtures/types.go

test-folder:
	./struct2oas -leftpad "  " -source fixtures

clean:
	rm -f *.yml
