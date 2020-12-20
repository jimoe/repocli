.PHONY: init build

BIN?=~/bin

build:
ifeq (found, $(shell test -d ${BIN} && echo found))
	@echo Building cli at ${BIN}
	@go build -o ${BIN}/repocli cmd/main.go
else
	$(error "Folder '${BIN}' does not exist. You may specify path: 'make init BIN=/some/path'")
endif

init: build
# thanks!: https://stackoverflow.com/a/25668869
ifneq (, $(shell which repocli))
	@repocli config init ${BIN}
else
	$(error "The cli was not found in PATH. Make sure ${BIN} is in PATH and run 'make init' again")
endif
