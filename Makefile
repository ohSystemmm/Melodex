APP_NAME=melodex

all: build

build:
	go build -o $(APP_NAME) .

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

deps:
	go mod tidy

fmt:
	go fmt ./...

test:
	go test ./...

help:
	@echo "Makefile commands:"
	@echo "  make build   - Build the project"
	@echo "  make run     - Build and run the project"
	@echo "  make clean   - Remove binary files"
	@echo "  make deps    - Install/update dependencies"
	@echo "  make fmt     - Format the code"
	@echo "  make test    - Run tests"
