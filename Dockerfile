FROM golang:latest

WORKDIR /app

COPY . .

CMD ["go", "run", "/app/main.go"]