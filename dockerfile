FROM golang:1.22.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin ./cmd/http

EXPOSE 8080

ENTRYPOINT ["/app/bin"]