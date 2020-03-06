.SUFFIXES:  # erase all the builtin rules

all: test install

.PHONY: test
test: fmt
	go test

.PHONY: fmt
fmt:
	go fmt .
	~/go/bin/goimports -w .

install: test
	go install ./okpw

clean:

