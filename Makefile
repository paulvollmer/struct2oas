test:
	go build
	./struct2oas -source fixtures/model1.go
