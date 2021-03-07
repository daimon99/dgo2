.PHONY: clean clean-test clean-pyc clean-build docs help
.DEFAULT_GOAL := help

help:           ## Show this help.
	# @fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | sort | awk -F ':.*?## ' 'NF==2 {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

build:  ## build
	go build -o dgo2
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o release/dgo2_linux main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o release/dgo2_mac main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o release/dgo2_arm main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o release/dgo2_win main.go

release: build ## release
	for i in release/dgo2_*;do ./dgo2 upload $$i; done;

