all:
	go build -o push-swap ./cmd/push-swap
	go build -o checker ./cmd/checker

test:
	go test ./...

clean:
	rm -f push-swap checker

re: clean all

.PHONY: all test clean re
