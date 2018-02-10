# Go parameters
GOCMD=go
BIN_DIR=./bin
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=wassit
CROSS_PLATFORMS=linux windows darwin

all: test build
build:
	# Build binary
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v
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
deps:
	# Install dependencires
	$(GOGET) github.com/golang/dep
	dep ensure


# Cross compilation
build-all:
	GOOS=darwin GOARCH=amd64 go build -v -o ${BIN_DIR}/k-update-darwin-amd64
	GOOS=linux GOARCH=amd64 go build -v -o ${BIN_DIR}/k-update-linux-amd64
	GOOS=windows GOARCH=amd64 go build -v -o ${BIN_DIR}/k-update-windows-amd64.exe
