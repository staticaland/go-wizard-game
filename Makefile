all:
	go run .
test:
	go test -v ./...
build:
	go build -o myapp
clean:
	go clean
fmt:
	go fmt ./...
vet:
	go vet ./...
	
.PHONY: all test build clean fmt vet