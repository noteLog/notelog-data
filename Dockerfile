FROM golang:latest as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CG0_ENABLED=0 GOOS=linux go build -mod=readonly -v

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/notelog-data /notelog-data

CMD ["./notelog-data"]
