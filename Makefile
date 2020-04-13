all: test

.PHONY: test
test:
	go test -cover -v