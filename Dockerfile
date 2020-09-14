FROM golang:1.12-alpine

RUN apk add --no-cache git

ENV GO111MODULE=on
WORKDIR /app/hello-fresh

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o ./out/hello-fresh .

EXPOSE 9000
CMD ["./out/hello-fresh"]

