FROM golang:1.16-alpine

WORKDIR /app
COPY client ./client
WORKDIR /app/client

CMD ["go", "run", "main.go"]