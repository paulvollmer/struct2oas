test:
	go build
	./struct2oas -leftpad "  " -source fixtures/model1.go
	./struct2oas -leftpad "  " -source fixtures/model2.go
