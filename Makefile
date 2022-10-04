.PHONY: install run
.DEFAULT_GOAL:= run

GOBIN=$(shell go env GOBIN)

install:
		go install godb/cmd/godb

run: install
		$(GOBIN)/godb
