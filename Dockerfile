FROM golang:1.21.5-alpine3.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o ./bin/server ./cmd/kasper/main.go

EXPOSE 8080

CMD ["./bin/server"]