GOCMD = go
GOBUILD = $(GOCMD) build
GO_BINARY_OUTPUT = ./bin/go-quick-link

GO_PROJECT_LINK = github.com/harrydrippin/go-quick-link

all: build run
build:
	$(GOBUILD) -o $(GO_BINARY_OUTPUT) github.com/harrydrippin/go-quick-link

run:
	$(GO_BINARY_OUTPUT)