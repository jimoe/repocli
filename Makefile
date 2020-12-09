.PHONY: build

BIN?=~/bin

#build:
#	@if [ -d ${BIN} ]; then \
#  	go build -o ${BIN}/repocli cmd/main.go; \
#	else \
#		echo "folder '${BIN}' does not exist."; \
#		echo "you may give another path: 'make init BIN=/some/path'"; \
#		exit 1; \
#  fi

build:
ifeq (found, $(shell test -d ${BIN} && echo found))
	@echo building...
	@go build -o ${BIN}/repocli cmd/main.go
else
	$(error "Folder '${BIN}' does not exist. You may give another path: 'make init BIN=/some/path'")
endif
