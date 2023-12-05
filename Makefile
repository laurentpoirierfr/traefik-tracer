.PHONY: test

GO111MODULE=on

default: test

test:
	go test -v -cover ./...

clean:
	rm -rf ./vendor