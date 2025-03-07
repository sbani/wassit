# Go parameters
GOCMD=go
BIN_DIR=./bin
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=wassit

all: test build
build:
	# Build binary
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v
update:
	go get -u
	go mod tidy
test:
	# Run tests
	$(GOTEST) -v ./...
clean:
	# Clean stuff
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	# Run binary
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v ./...
	./$(BINARY_NAME)


# Cross compilation
build-all:
	GOOS=darwin GOARCH=amd64 go build -v -o ${BIN_DIR}/$(BINARY_NAME)-darwin-amd64
	GOOS=linux GOARCH=amd64 go build -v -o ${BIN_DIR}/$(BINARY_NAME)-linux-amd64
	GOOS=windows GOARCH=amd64 go build -v -o ${BIN_DIR}/$(BINARY_NAME)-windows-amd64.exe
