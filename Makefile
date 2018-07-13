init:
	go get -u github.com/golang/dep/cmd/dep
	dep init;

vendor: init
	dep ensure -update

test:
	go test -v ./... -cover
