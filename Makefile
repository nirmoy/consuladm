SHELL := /usr/bin/env bash
CWD := $(shell pwd)
BIN := consuladm

.PHONY: clean

all: $(BIN)

$(BIN):
	go build -o $(BIN) main.go

clean:
	rm -f $(BIN)
