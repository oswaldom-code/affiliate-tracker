# syntax=docker/dockerfile:1

FROM golang:1.18
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=0 go build -o bin/go-template main.go

ENTRYPOINT ["/app/bin/go-template", "server"]
EXPOSE 9000
