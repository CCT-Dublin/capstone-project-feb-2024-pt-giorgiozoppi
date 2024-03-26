.PHONY: all build run clean
all: clean build
build:
	go build -o downloader

run: build
	./downloader

clean:
	rm -f downloader