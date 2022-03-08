ifeq ($(GOPATH),)
GOPATH := $(HOME)/go
endif

all: test lint vet build

build: browser service

browser:
	@echo "*** building $@"
	@cd cmd/$@ && go build -o ../../bin/$@ --trimpath -tags osusergo,netgo -ldflags="-s -w"

service:
	@echo "*** building $@"
	@cd cmd/$@ && go build -o ../../bin/$@ --trimpath -tags osusergo,netgo -ldflags="-s -w"

test:
	@echo "*** $@"
	@go test ./...

vet:
	@echo "*** $@"
	@go vet ./...

lint:
	@echo "*** $@"
	@revive ./... 

clean:
	@rm -rf bin
