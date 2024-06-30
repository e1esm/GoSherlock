BINARY_NAME = gosherlock

.PHONY: clean
clean:
	rm -rf bin/$(BINARY_NAME)

.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) ./src/$(BINARY_NAME)/cmd/.