SHELL := /usr/bin/env bash
CWD := $(shell pwd)
BIN := consuladm

.PHONY: clean

all: $(BIN)

$(BIN):
	GO111MODULE=on go build -o $(BIN) main.go

clean:
	rm -f $(BIN)
