## Weather bot app for telegram written in Golang


### Prerequisites
 ...

### Building & running locally
Populate the *.env* file with both your Telegram Bot Token & your Open Weather Map API Key
1. **Using **_make_****


**Init the variables**
```shell
make init
```

**Compile**
```shell
make build
```
**Run the unit tests**
```shell
make test
```
**Running the Bot**
```shell
make run
```
**List commands**
```shell
make help
```

2. **Using go **_CLI_****
```shell
export $(grep -v '^#' .env | xargs -d '\n')
```
**Compile**
```shell
go build -o bot
```
**Run the unit tests**
```shell
go test
```
**Running the Bot**
```shell
./bot
```
**Output:**

### Building the docker image and running it
Build the Docker image

```shell
docker build -t bot .
```

Run the built docker image, please make sure to provide the environement variables file *.env*.
```shell
docker run -it --rm --env-file=.env --name mybot bot
```
