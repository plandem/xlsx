init:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint
	dep init;

vendor: init
	dep ensure -update

lint:
	golint ./...

bench:
	go test -run=^$$ -bench=Libs

bench-sheet:
	go test -run=^$$ -bench=Spreadsheet

test:
	go test -v ./... -cover
