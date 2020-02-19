FROM golang:latest
RUN mkdir $GOPATH/src/notelog-data
WORKDIR $GOPATH/src/notelog-data
COPY . .
RUN go mod download
RUN go build .
CMD ["./notelog-data"]
