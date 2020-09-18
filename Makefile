BIN = mackerel-plugin-qdisc

all: clean  build

build:
	go build -o bin/$(BIN) main.go

clean:
	rm -rf bin

.PHONY: build clean
