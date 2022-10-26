FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/vincentconace/golang-docker

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]