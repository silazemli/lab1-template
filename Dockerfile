FROM golang:1.23.0

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY .. ./

RUN go build -o ./build/main ./cmd/main.go

EXPOSE 8080
ENTRYPOINT ./build/main
