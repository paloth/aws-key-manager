# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' ./...)

DATE=$(shell date -u "+%a %b %d %T %Y")

LDFLAGS = -s -w
LDFLAGS += -X "github.com/b4nst/turbogit/cmd.BuildDate=$(DATE)"
LDFLAGS += -X "github.com/b4nst/turbogit/cmd.Commit=$(TUG_COMMIT)"
LDFLAGS += -X "github.com/b4nst/turbogit/cmd.Version=$(TUG_VERSION)"

bin/akm: $(BUILD_FILES)
	$(GOBUILD) -trimpath -o "$@" -ldflags='$(LDFLAGS)' ./main.go 
build: bin/akm
.PHONY: build

test: $(BUILD_FILES)
	$(GOTEST) ./...  -coverprofile c.out
.PHONY: test

doc:
	$(GORUN) scripts/gen-doc.go
.PHONY: doc

clean:
	rm -rf bin dist c.out
.PHONY: clean
