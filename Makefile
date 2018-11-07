 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=notify
BINARY_THIS=$(BINARY_NAME)_this
BINARY_LINUX=$(BINARY_NAME)_unix
BINARY_DARWIN=$(BINARY_NAME)_darwin
BINARY_WINDOWS=$(BINARY_NAME)_windows.exe

all: build

build: build-this build-linux build-darwin build-windows
	

test: 
		$(GOTEST) -v ./...

clean: 
		$(GOCLEAN)
		rm -f $(BINARY_THIS) $(BINARY_LINUX) $(BINARY_DARWIN) $(BINARY_WINDOWS)

run: build-this
		./$(BINARY_THIS)

build-this:
	$(GOBUILD) -o $(BINARY_THIS)
	
# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX)
build-darwin:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN)
build-windows:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WINDOWS)
