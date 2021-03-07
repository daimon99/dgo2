.PHONY: clean clean-test clean-pyc clean-build docs help
.DEFAULT_GOAL := help

help:           ## Show this help.
	# @fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | sort | awk -F ':.*?## ' 'NF==2 {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

build:  ## build
	go build
