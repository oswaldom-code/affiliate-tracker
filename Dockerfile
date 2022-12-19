# syntax=docker/dockerfile:1

FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=0 go build -o bin/affiliate-tracker main.go

ENTRYPOINT ["/app/bin/affiliate-tracker", "server"]
EXPOSE 3000
