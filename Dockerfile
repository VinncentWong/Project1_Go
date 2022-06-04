FROM golang:alpine3.15

# Make a directory in our docker image
WORKDIR "/app"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o "/executablefile"

CMD ["/executablefile"]