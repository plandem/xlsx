init:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint
	dep init;

vendor: init
	dep ensure -update

lint:
	goreportcard-cli -v

bench:
	go test -run=^$$ -bench=Libs

bench-sheet:
	go test -run=^$$ -bench=Spreadsheet

test:
	go test -v ./... -cover

test-docs:
	go test -v ./docs

docs-vendor:
	yarn --cwd ./docs install

docs-build: docs-vendor
	yarn --cwd ./docs build

docs-dev: docs-vendor
	yarn --cwd ./docs dev

