FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]
