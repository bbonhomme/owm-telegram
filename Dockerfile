# The base go-image
FROM golang:latest AS builder

# Create a directory for the app
RUN mkdir /app
 
# Set working directory
WORKDIR /app

#  Go modules on
ENV GO111MODULE=on


# Get the modules & dep. ready to download
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy all files from the current directory to the app directory
COPY . /app

# Run command as described:
# go build will build an executable file named bot in the current directory
#RUN go build -v -o bot .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bot .

# Use a lightweight image to run the bot
#FROM ubuntu:bionic
#COPY --from=builder /app/bot /app/bot

# Run the server executable
ENTRYPOINT [ "/app/bot" ]