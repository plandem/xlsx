init:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint
	dep init;

vendor: init
	dep ensure -update

lint:
	goreportcard-cli -v

test:
	go test -v ./... -cover

test-docs:
	go test -v ./docs

docs-vendor:
	yarn --cwd ./docs install

docs-build: docs-vendor
	yarn --cwd ./docs build

docs-publish: docs-build
	cd ./docs/src/.vuepress/dist; git init; git add -A; git commit -m 'deploy'; git push -f git@github.com:plandem/xlsx.git master:gh-pages;

docs-dev: docs-vendor
	yarn --cwd ./docs dev

