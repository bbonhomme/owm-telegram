help:
	echo "make build: to compile the source code\nmake init: to export the variables defined in .env file \nmake run: to run the compile code\nmake test: to run the unit tests"

IMAGE := bbonhomme/owm-telegram


init:
	export $(grep -v '^#' .env | xargs -d '\n')

build:
	go build -o bot

test:
	go test

run:
	./bot

all: help

build-image:
	docker build -t $(IMAGE) .

run-image:
	docker run -it -d  --rm -e "OWM_API_KEY=$OWM_API_KEY" -e "TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN"  --name bot-image-run $(IMAGE)

push-image:
	docker push $(IMAGE)

.PHONY: image push-image test