BINARY=algosort

build:
	go build -o bin/$(BINARY) main.go

run:
	go run main.go

test:
	go test ./sorting
