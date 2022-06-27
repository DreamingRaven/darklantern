.PHONY: all
all: tidy style test clean

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: style
style:
	gofmt -s -w -l .

.PHONY: test
test:
	./lazy-test.sh
	# go test $(go list ./... | grep -v /vendor/)

.PHONY: clean
clean:
