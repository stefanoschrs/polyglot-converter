#!make

upx := $(shell which upx)
out := "polyglot-converter"

build:
	go build -ldflags "-s -w" -o $(out) .
ifdef upx
	upx $(out)
endif
