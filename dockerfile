FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin ./cmd/http

EXPOSE 8080

ENTRYPOINT [ "/app/bin" ]