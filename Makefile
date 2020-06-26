help:
	echo "make build: to compile the source code\nmake init: to export the variables defined in .env file \nmake run: to run the compile code\nmake test: to run the unit tests"

init:
	export $(grep -v '^#' .env | xargs -d '\n')

build:
	go build -o bot

test:
	go test

run:
	./bot

all: help
