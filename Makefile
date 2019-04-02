test:
	go build
	./struct2oas -leftpad "  " -source fixtures/model1.go
