FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/rpc/opencc

CMD ["go", "run", "."]
