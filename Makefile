## personal dev is on macOS, but this should be
## supporting Linux as well
## gfind can be installed with `brew install findutils`
UNAME := $(shell uname)
ifeq ($(UNAME),Linux)
	FIND_BIN=find
endif
ifeq ($(UNAME),Darwin)
	FIND_BIN=gfind
endif

.SHELLFLAGS = -ec
.ONESHELL:

build: ./bin/api-server

./bin/api-server: $(shell $(FIND_BIN) -name "*.go")
	go build -o ./bin/api-server ./cmd/api-server

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: image
image:
	docker build -f ./cmd/api-server/Dockerfile -t latest .

.PHONY: run
run:
	./bin/api-server
